package img

import (
	"encoding/base64"
	"fmt"
	"mime"
	"os"
	"path/filepath"
)

func LocalImageToDataURL(imagePath string) (string, error) {
	// Guess the MIME type of the image based on the file extension
	ext := filepath.Ext(imagePath)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "application/octet-stream" // Default MIME type if none is found
	}

	// Read and encode the image file
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		return "", err
	}

	base64EncodedData := base64.StdEncoding.EncodeToString(imageData)

	// Construct the data URL
	dataURL := fmt.Sprintf("data:%s;base64,%s", mimeType, base64EncodedData)
	return dataURL, nil
}

func DecodeBase642Img(base64Str, imgPath string) error {
	// Decode the base64 string
	imageData, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return err
	}

	// Write the image data to a file
	err = os.WriteFile(imgPath, imageData, 0644)
	if err != nil {
		return err
	}

	return nil
}
