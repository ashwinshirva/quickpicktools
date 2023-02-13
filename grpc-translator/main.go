package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "github.com/quickpicktools/grpc-translator/proto/grpc-translator"
	"github.com/sirupsen/logrus"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedFrontendServer
}

func (s *server) Homepage(ctx context.Context, request *pb.HomepageRequest) (*pb.HomepageRespone, error) {
	logrus.Info("request.Message: ", request.GetMessage())
	name := request.Message
	response := &pb.HomepageRespone{
		Message: "Hello " + name,
	}
	return response, nil
}

func main() {
	// // Create a listener on TCP port
	// lis, err := net.Listen("tcp", ":8080")
	// if err != nil {
	// 	log.Fatalln("Failed to listen:", err)
	// }

	// // Create a gRPC server object
	// s := grpc.NewServer()
	// // Attach the Greeter service to the server
	// pb.RegisterFrontendServer(s, &server{})
	// // Serve gRPC server
	// log.Println("Serving gRPC on connection ")
	// go func() {
	// 	log.Fatalln(s.Serve(lis))
	// }()

	// // Create a client connection to the gRPC server we just started
	// conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalln("Failed to dial server:", err)
	// }
	// defer conn.Close()

	// mux := runtime.NewServeMux()

	// // Register Greeter
	// err = pb.RegisterFrontendHandler(context.Background(), mux, conn)
	// if err != nil {
	// 	log.Fatalln("Failed to register gateway:", err)
	// }

	// gwServer := &http.Server{
	// 	Addr:    "",
	// 	Handler: mux,
	// }

	// log.Println("Serving gRPC-Gateway on connection")
	// log.Fatalln(gwServer.ListenAndServe())

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterFrontendServer(s, &server{})
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = pb.RegisterFrontendHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())

}
