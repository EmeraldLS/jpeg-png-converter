package conversion

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
)

func PngToJpeg(imageBytes []byte, outputFileName string) (string, error) {
	imgType := http.DetectContentType(imageBytes)
	if imgType == "image/png" {
		jpegBytes, err := runPngConversionToJpeg(imgType, imageBytes)
		if err != nil {
			return "", err
		}
		filePath := fmt.Sprintf("./out/%v.jpg", outputFileName)
		err = os.WriteFile(filePath, jpegBytes, 0644)
		if err != nil {
			return "", err
		}
		return "image converted successfully", nil
	}
	return "", fmt.Errorf("unable to convert png file to jpeg")
}

func runPngConversionToJpeg(imgType string, imageBytes []byte) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, nil); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
