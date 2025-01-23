package img

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

type Data struct {
	Total int `json:"total"`
	Start int `json:"start"`
	Data  []struct {
		ServerID  string `json:"serverId"`
		TimeStamp int64  `json:"timeStamp"`
		MapID     string `json:"mapId"`
		MapData   string `json:"mapData"`
	} `json:"data"`
}

func TestDecodeBase642Img(t *testing.T) {
	t.Log("TestDecodeBase642Img")
	imgPath := "test.png"
	base64Str := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkYAAAAAYAAjCB0C8AAAAASUVORK5CYII="
	ret := DecodeBase642Img(base64Str, imgPath)
	if !ret {
		t.Error("DecodeBase642Img() failed")
		return
	}
	t.Log("DecodeBase642Img() succeeded")
}
func TestCatchImgCode(t *testing.T) {
	// Read the JSON file
	file, err := os.Open("/Users/admin/Downloads/response.json")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Read the file content
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	// Unmarshal the JSON data
	var data Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatalf("failed to unmarshal JSON: %s", err)
	}

	// Create or open the output file
	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("failed to create output file: %s", err)
	}
	defer outputFile.Close()

	// Process each mapData and add the prefix
	for _, item := range data.Data {
		prefixedMapData := "data:image/png;base64," + item.MapData
		_, err := outputFile.WriteString(prefixedMapData + "\n")
		if err != nil {
			log.Fatalf("failed to write to output file: %s", err)
		}
	}

	fmt.Println("Data has been written to output.txt")
}

func TestAvgSplitImg(t *testing.T) {
	imgPath := "/Users/admin/Downloads/1.png"
	targetDir := "/Users/admin/Downloads/temp/"
	count := 9
	list, err := AvgSplitImg(count, imgPath, targetDir)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("AvgSplitImg() succeeded", list)
	return
}
