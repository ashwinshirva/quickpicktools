package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ashwinshirva/quickpicktools/proto-gen/go/grpc-translator"
	pbToJpg "github.com/ashwinshirva/quickpicktools/proto-gen/go/to-jpg-service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
)

var toJPGClientConn *grpc.ClientConn

func RegisterServices() {

	/* 	//go shutdownhook()

	   	// Create a listener on TCP port
	   	lis, err := net.Listen("tcp", ":8082")
	   	if err != nil {
	   		log.Fatalln("Failed to listen:", err)
	   	}

	   	// Create a gRPC server object
	   	s := grpc.NewServer()
	   	// Attach the Frontend service to the server
	   	pb.RegisterFrontendServer(s, &server{})
	   	// Serve gRPC server
	   	log.Println("Serving gRPC on 0.0.0.0:8082")
	   	go func() {
	   		log.Fatalln(s.Serve(lis))
	   	}()

	   	test() */

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		//"0.0.0.0:8089",
		"qpt-to-jpg-service:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.MaxCallRecvMsgSize(100*1024*1024), //math.MaxInt32),
		// grpc.MaxCallSendMsgSize(100*1024*1024), //math.MaxInt32),
		//grpc.MaxMsgSize(100*1024*1024),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	SetToJPGClientConn(conn)

	gwmux := runtime.NewServeMux()

	// Attachment upload from http/s handled manually
	// This is needed because grpc-gateway does not support multipart uploads(binary uploads)
	if err := gwmux.HandlePath("POST", "/to-jpg/png-to-jpg", handleBinaryFileUpload); err != nil {
		panic(err)
	}

	// Register Greeter
	err = pb.RegisterFrontendHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// Register ToJPGService
	// Register Greeter
	/* err = pbToJpg.RegisterToJpgServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register ToJPGService to gateway:", err)
	} */

	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8080")
	log.Fatalln(gwServer.ListenAndServe())
}

func SetToJPGClientConn(conn *grpc.ClientConn) {
	toJPGClientConn = conn
}

func handleBinaryFileUpload(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Info("handleBinaryFileUpload::handleBinaryFileUpload called...")
	err := r.ParseForm()
	if err != nil {
		log.Error("handleBinaryFileUpload::Error parsing form: ", err)
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	f, header, err := r.FormFile("image")
	if err != nil {
		log.Error("handleBinaryFileUpload::Error getting image from form: ", err)
		http.Error(w, fmt.Sprintf("failed to get file 'attachment': %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer f.Close()

	//
	// Now do something with the io.Reader in `f`, i.e. read it into a buffer or stream it to a gRPC client side stream.
	// Also `header` will contain the filename, size etc of the original file.
	//
	// Get type, identifier from params
	/* ofType, ok := params["type"]
	if !ok {
		writeErr(http.StatusBadRequest, "Missing 'type' param", w)
		return
	}

	identifier, ok := params["identifier"]
	if !ok {
		writeErr(http.StatusBadRequest, "Missing 'identifier' param", w)
		return
	}

	err = rh.store.Attach(ofType, identifier, header.Filename, f)
	if err != nil {
		writeErr(http.StatusInternalServerError, err.Error(), w)
		return
	} */
	pngToJPGResp, pngToJPGErr := DialPngToJpg(header.Filename, f)
	if pngToJPGErr != nil {
		log.Error("handleBinaryFileUpload::Error from server: ", pngToJPGErr)
	}

	//w.WriteHeader(http.StatusOK)
	// Set the content type to JSON and return the REST response
	/* w.Header().Set("Content-Type", "image/jpg")
	json.NewEncoder(w).Encode(pngToJPGResp) */

	//=========================================
	// Create a new bytes.Buffer object
	buf := new(bytes.Buffer)

	// Write the byte slice to the buffer
	_, err = buf.Write(pngToJPGResp.GetConvertedImage().GetData())
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
	}

	// Get the image data
	imgData, err := ioutil.ReadAll(buf)
	if err != nil {
		http.Error(w, "Error reading image data", http.StatusInternalServerError)
		return
	}

	// Set the content type header
	w.Header().Set("Content-Type", "application/json")

	// Create the response struct
	response := ImageResponse{
		Data: map[string]string{
			"name": string(pngToJPGResp.GetConvertedImage().GetMetadata().GetName()),
		},
		Image: imgData,
	}

	// Encode the response as JSON
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Write the JSON data and image data to the response body
	w.Write(jsonBytes)
}

type ImageResponse struct {
	Data  interface{} `json:"data"`
	Image []byte      `json:"image"`
}

/* func CreateMultipartFile(imgBuf []byte, imgName string) (multipart.File, error) {
	// create a new multipart buffer
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	// add the byte slice as a file field
	part, err := writer.CreateFormFile("imgfile", imgName)
	if err != nil {
		// handle error
	}
	_, err = part.Write(imgBuf)
	if err != nil {
		// handle error
	}

	// close the multipart writer
	err = writer.Close()
	if err != nil {
		// handle error
	}
	return buf, nil
} */

func DialPngToJpg(fileName string, file multipart.File) (*pbToJpg.PngToJpgResp, error) {
	log.Info("DialPngToJpg:: DialPngToJpg() called...")

	fileBuf := bytes.NewBuffer(nil)
	if _, err := io.Copy(fileBuf, file); err != nil {
		log.Error("DialPngToJpg::Error writing PNG image to buffer: ", err)
		return nil, err
	}

	client := pbToJpg.NewToJpgServiceClient(toJPGClientConn)
	toJPGResp, toJPGErr := client.PngToJpg(context.Background(), &pbToJpg.PngToJpgReq{ImageUrl: fileName, Image: &pbToJpg.Image{Metadata: &pbToJpg.ImageMetadata{Name: fileName}, Data: fileBuf.Bytes()}})
	if toJPGErr != nil {
		log.Error("DialPngToJpg::Error connecting to worker service: ", toJPGErr)
		return nil, toJPGErr
	}
	//log.Info("DialPngToJpg::Worker service response: ", toJPGResp)
	return toJPGResp, nil
}

/* func test(fileName string) {
	log.Info("Test() called...")
	conn, err := grpc.Dial("qpt-to-jpg-service:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("Error dialling worker: ", err)
		return
	}
	client := pbToJpg.NewToJpgServiceClient(conn)
	resp, err2 := client.PngToJpg(context.Background(), &pbToJpg.PngToJpgReq{ImageUrl: fileName, Image: &pbToJpg.Image{Metadata: &pbToJpg.ImageMetadata{Name: fileName}, Data: imageData})
	if err2 != nil {
		log.Error("Error connecting to worker service: ", err2)
		return
	}
	log.Info("Worker service response: ", resp)
} */

func shutdownhook() {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)
	<-channel
	os.Exit(0)
}
