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

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pbToJpg "github.com/ashwinshirva/quickpicktools/proto-gen/go/to-jpg-service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
)

var toJPGClientConn *grpc.ClientConn

func RegisterServices() {
	grpcGatewayPort := "8080"
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"qpt-to-jpg-service:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
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

	gwServer := &http.Server{
		Addr:    ":" + grpcGatewayPort,
		Handler: gwmux,
	}

	log.Info("Serving gRPC-Gateway on http://0.0.0.0:" + grpcGatewayPort)
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

	pngToJPGResp, pngToJPGErr := DialPngToJpg(header.Filename, f)
	if pngToJPGErr != nil {
		log.Error("handleBinaryFileUpload::Error from server: ", pngToJPGErr)
	}

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
	return toJPGResp, nil
}
