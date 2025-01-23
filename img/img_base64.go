package img

import (
	"encoding/base64"
	"fmt"
	"github.com/otkinlife/go_tools/file_tools"
	"math/rand"
	"mime"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
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

func DecodeBase642Img(path string, base64ImgStr string) bool {
	b, _ := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, base64ImgStr)
	if !b {
		return false
	}
	re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
	allData := re.FindAllSubmatch([]byte(base64ImgStr), 2)
	fileType := string(allData[0][1]) //png ，jpeg 后缀获取

	base64Str := re.ReplaceAllString(base64ImgStr, "")

	date := time.Now().Format("20060102")
	if ok := file_tools.IsFileExist(path + "/" + date); !ok {
		os.Mkdir(path+"/"+date, 0666)
	}

	var file = path + "/" + date + "/" + strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+100000) + "." + fileType
	byte, _ := base64.StdEncoding.DecodeString(base64Str)

	err := os.WriteFile(file, byte, 0666)
	if err != nil {
		return false
	}
	return true
}
