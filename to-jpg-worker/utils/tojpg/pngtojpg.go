package tojpg

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"

	imgLib "github.com/ashwinshirva/quickpicktools/qpt-lib/image"
	log "github.com/sirupsen/logrus"
)

// PNGToJPG converts a PNG image to JPG image
func PNGToJPG(imageName string, imageData []byte) error {
	// byte slice to bytes.Reader, which implements the io.Reader interface
	reader := bytes.NewReader(imageData)

	// Decode PNG image
	pngImage, err := png.Decode(reader)
	if err != nil {
		log.Error("PNGToPDF::Error decoding png image: ", err)
		return err
	}
	return convertToJPG(imageName, pngImage)
}

// convertToJPG converts the given source image to JPG image
func convertToJPG(imageName string, imageSrc image.Image) error {
	// Create a new Image with the same dimension of PNG image
	newImg := image.NewRGBA(imageSrc.Bounds())

	// Using white background to replace PNG's transparent background
	// To change it to other color use a new color.RGBA{}
	// and use image.NewUniform(color.RGBA{<fill in color>}) function
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// Paste PNG image OVER to newImage
	draw.Draw(newImg, newImg.Bounds(), imageSrc, imageSrc.Bounds().Min, draw.Over)

	// Get converted image name
	convImgName, convErr := imgLib.ConvertedImageName(imageName, "jpg")
	if convErr != nil {
		log.Error("convertToJPG::Error getting converted image name: ", convErr)
		return convErr
	}

	// Create new out JPEG file
	jpgImgFile, err := os.Create(convImgName)
	if err != nil {
		log.Error("convertToJPG::Error creating JPG image file: ", err)
		return err
	}
	defer jpgImgFile.Close()

	// Convert newImage to JPEG encoded byte and save to jpgImgFile with quality = 90
	// err = jpeg.Encode(jpgImgFile, newImg, nil) -- use nil if ignore quality options
	var opt jpeg.Options
	opt.Quality = 90
	err = jpeg.Encode(jpgImgFile, newImg, &opt)
	if err != nil {
		log.Error("convertToJPG::Error encoding to JPG image: ", err)
		return err
	}
	return nil
}
