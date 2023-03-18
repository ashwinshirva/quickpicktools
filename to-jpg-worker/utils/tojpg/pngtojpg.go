package tojpg

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"

	log "github.com/sirupsen/logrus"
)

// PNGToJPG converts a PNG image to JPG image
func PNGToJPG(imageName string, imageData []byte) error {
	/* // Open png image in the given path
	pngFile, err := os.Open(pathToImage)
	if err != nil {
		log.Error("PNGToPDF::Error occured opening png image: ", err)
		return err
	}
	defer pngFile.Close() */

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

func convertToJPG(imageName string, imageSrc image.Image) error {
	// Create a new Image with the same dimension of PNG image
	newImg := image.NewRGBA(imageSrc.Bounds())

	// we will use white background to replace PNG's transparent background
	// you can change it to whichever color you want with
	// a new color.RGBA{} and use image.NewUniform(color.RGBA{<fill in color>}) function

	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// paste PNG image OVER to newImage
	draw.Draw(newImg, newImg.Bounds(), imageSrc, imageSrc.Bounds().Min, draw.Over)

	// create new out JPEG file
	jpgImgFile, err := os.Create(imageName + "_converted.jpg")

	if err != nil {
		fmt.Println("Cannot create JPEG-file.jpg !")
		fmt.Println(err)
		return err
	}

	defer jpgImgFile.Close()

	var opt jpeg.Options
	opt.Quality = 80

	// convert newImage to JPEG encoded byte and save to jpgImgFile
	// with quality = 80
	err = jpeg.Encode(jpgImgFile, newImg, &opt)

	//err = jpeg.Encode(jpgImgFile, newImg, nil) -- use nil if ignore quality options

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
