package services

import (
	"context"

	"github.com/ashwinshirva/quickpicktools/to-jpg-worker/utils/tojpg"
	log "github.com/sirupsen/logrus"

	pb "github.com/ashwinshirva/quickpicktools/proto-gen/go/to-jpg-service"
)

type ToJPGService struct {
	pb.UnimplementedToJpgServiceServer
}

func (s *ToJPGService) PngToJpg(ctx context.Context, request *pb.PngToJpgReq) (*pb.PngToJpgResp, error) {
	imageName := request.GetImage().GetMetadata().GetName()
	log.Infof("ToJPGService::PngToJpg::Converting image %s into JPG", imageName)
	resp := &pb.PngToJpgResp{Status: &pb.ResponseStatus{Status: pb.Status_SUCCESS}}
	// TODO: This hardcoded value needs to be removed
	//err := tojpg.PNGToJPG("./test.png")
	err := tojpg.PNGToJPG(imageName, request.GetImage().GetData())
	if err != nil {
		log.Error("ToJPGService::PngToJpg::Error converting image: ", err)
		return nil, err
	}
	return resp, nil
}
