package services

import (
	"context"

	"github.com/sirupsen/logrus"

	pb "github.com/ashwinshirva/quickpicktools/proto-gen/go/grpc-translator"
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
