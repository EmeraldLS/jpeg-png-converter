package conversion

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
)

func ToPng(imageBytes []byte, outputFileName string) error {
	imgType := http.DetectContentType(imageBytes)
	if imgType == "image/jpeg" || imgType == "image/jpg" {
		pngBytes, err := runJpgConversionToPng(imageBytes)
		if err != nil {
			return err
		}
		filePath := fmt.Sprintf("./out/%v.png", outputFileName)
		if err = os.WriteFile(filePath, pngBytes, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("image type isn't jpeg")
}

func runJpgConversionToPng(imageBytes []byte) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	if err = png.Encode(buf, img); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
