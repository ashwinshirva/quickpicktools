package services

import (
	"context"

	"github.com/ashwinshirva/quickpicktools/to-jpg-worker/utils/tojpg"
	log "github.com/sirupsen/logrus"

	pb "github.com/ashwinshirva/quickpicktools/proto-gen/go/to-jpg-service"
	"github.com/ashwinshirva/quickpicktools/qpt-lib/image"
	imgLib "github.com/ashwinshirva/quickpicktools/qpt-lib/image"
)

// ToJPGService represents ToJPGService type
type ToJPGService struct {
	pb.UnimplementedToJpgServiceServer
}

// ToJPGService rpc method to convert PNG to JPG type
func (s *ToJPGService) PngToJpg(ctx context.Context, request *pb.PngToJpgReq) (*pb.PngToJpgResp, error) {
	imageName := request.GetImage().GetMetadata().GetName()

	log.Infof("ToJPGService::PngToJpg::Converting image %s into JPG", imageName)

	resp := &pb.PngToJpgResp{Status: &pb.ResponseStatus{Status: pb.Status_SUCCESS}}

	// Convert PNG to JPG
	err := tojpg.PNGToJPG(imageName, request.GetImage().GetData())
	if err != nil {
		log.Error("ToJPGService::PngToJpg::Error converting image: ", err)
		return nil, err
	}

	// Construct converted image name
	var convImgName string
	convImgName, err = imgLib.ConvertedImageName(imageName, "jpg")
	if err != nil {
		log.Error("ToJPGService::PngToJpg::Error constructing converted image name: ", err)
		return nil, err
	}

	// Convert JPG image to byte stream
	convImgPath := "./" + convImgName
	var convImgBuf []byte
	convImgBuf, err = imgLib.ConvertImageToBytes(convImgPath)
	if err != nil {
		log.Error("ToJPGService::PngToJpg::Error decoding converted image to buffer: ", err)
		return nil, err
	}
	resp.ConvertedImage = &pb.Image{Metadata: &pb.ImageMetadata{Name: convImgName}, Data: convImgBuf}

	// Delete the converetd image from path(after the response is sent back to the user)
	defer image.DeleteImage(convImgPath)

	return resp, nil
}
