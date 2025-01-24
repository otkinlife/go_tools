package img

import (
	"encoding/base64"
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"regexp"
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

func DecodeBase642Img(base64Str string, path string) error {
	// Define a regular expression to match the base64 prefix
	re := regexp.MustCompile(`^data:image\/[a-zA-Z]+;base64,`)

	// Find the prefix
	prefix := re.FindString(base64Str)
	if prefix == "" {
		return fmt.Errorf("invalid base64 string")
	}

	// Remove the prefix from the base64 string
	base64Data := base64Str[len(prefix):]

	// Decode the base64 string
	imageData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return err
	}

	// Ensure the directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Write the image data to the file
	if err := os.WriteFile(path, imageData, 0644); err != nil {
		return err
	}

	return nil
}

func DecodeBase64(base64Str string) ([]byte, error) {
	// Define a regular expression to match the base64 prefix
	re := regexp.MustCompile(`^data:image\/[a-zA-Z]+;base64,`)

	// Find the prefix
	prefix := re.FindString(base64Str)
	if prefix == "" {
		return nil, fmt.Errorf("invalid base64 string")
	}

	// Remove the prefix from the base64 string
	base64Data := base64Str[len(prefix):]

	// Decode the base64 string
	imageData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, err
	}

	return imageData, nil
}
