package image

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// DeleteImage delete image in the specified path
func DeleteImage(imagePath string) {
	err := os.Remove(imagePath)
	if err != nil {
		log.Warning("DeleteImage::Error occured deleting image: ", err)
	}
}

// ConvertedImageName returns the converted image name
// Naming is as per the Quick Pick Tools naming format for converted files
// Naming format for converted files: original_image_name + _converted + .image_filetype
func ConvertedImageName(originalImgName, convertedImageFiletype string) (string, error) {
	// originalImgName empty check
	if originalImgName == "" {
		log.Error("GetConvertedImageName::Error image name is empty...")
		return "", errors.New("image name is empty")
	}

	splitStr := strings.Split(originalImgName, ".")

	return splitStr[0] + "_converted." + convertedImageFiletype, nil
}

// ConvertImageToBytes converts the given image into a byte stream
func ConvertImageToBytes(filePath string) ([]byte, error) {
	// Load the image file
	imgFile, err := os.Open(filePath)
	if err != nil {
		log.Error("ConvertImageToBytes::Error opening image file: ", err)
		return nil, err
	}
	defer imgFile.Close()

	fileBuf := bytes.NewBuffer(nil)
	if _, err := io.Copy(fileBuf, imgFile); err != nil {
		log.Error("ConvertImageToBytes::Error writing JPG image to buffer: ", err)
		return nil, err
	}
	return fileBuf.Bytes(), nil
}
