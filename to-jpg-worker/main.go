package main

import (
	"net"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"

	pb "github.com/ashwinshirva/quickpicktools/proto-gen/go/to-jpg-service"
	"github.com/ashwinshirva/quickpicktools/to-jpg-worker/constants"
	"github.com/ashwinshirva/quickpicktools/to-jpg-worker/services"
)

func main() {
	// To JPG server port
	port := ":" + constants.DefaultPort

	// Setup a new TCP listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Infof("Listening on %s", port)

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register ToJPGService service to gRPC server
	pb.RegisterToJpgServiceServer(grpcServer, &services.ToJPGService{})

	// Start gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
