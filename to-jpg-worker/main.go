package main

import (
	"net"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"

	pb "github.com/ashwinshirva/quickpicktools/proto-gen/go/to-jpg-service"
	"github.com/ashwinshirva/quickpicktools/to-jpg-worker/services"
)

func main() {
	//tojpg.PNGToJPG("./test.png")

	port := ":8080"

	// Setup a new TCP listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Infof("Listening on %s", port)

	// Create a new gRPC server
	grpcServer := grpc.NewServer(grpc.MaxMsgSize(10*1024*1024),
		//grpc.MaxCallRecvMsgSize(100*1024*1024),
		//grpc.MaxCallSendMsgSize(100*1024*1024),
		grpc.MaxRecvMsgSize(100*1024*1024),
		grpc.MaxSendMsgSize(100*1024*1024))

	// Register ToJPGService service to gRPC server
	pb.RegisterToJpgServiceServer(grpcServer, &services.ToJPGService{})

	// Start gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
