package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
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
	DialPngToJpg(header.Filename, f)

	w.WriteHeader(http.StatusOK)
}

func DialPngToJpg(fileName string, file multipart.File) error {
	log.Info("DialPngToJpg:: DialPngToJpg() called...")

	fileBuf := bytes.NewBuffer(nil)
	if _, err := io.Copy(fileBuf, file); err != nil {
		log.Error("DialPngToJpg::Error writing PNG image to buffer: ", err)
		return err
	}

	client := pbToJpg.NewToJpgServiceClient(toJPGClientConn)
	toJPGResp, toJPGErr := client.PngToJpg(context.Background(), &pbToJpg.PngToJpgReq{ImageUrl: fileName, Image: &pbToJpg.Image{Metadata: &pbToJpg.ImageMetadata{Name: fileName}, Data: fileBuf.Bytes()}})
	if toJPGErr != nil {
		log.Error("DialPngToJpg::Error connecting to worker service: ", toJPGErr)
		return toJPGErr
	}
	log.Info("DialPngToJpg::Worker service response: ", toJPGResp)
	return nil
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
