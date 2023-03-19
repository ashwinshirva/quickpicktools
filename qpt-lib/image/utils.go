package image

import (
	"errors"
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
