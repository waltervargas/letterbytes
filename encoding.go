package letterbytes

import (
	"fmt"
	"image/png"
	"os"

	"github.com/tjgq/sane"
	"golang.org/x/image/tiff"
)

func toTIFF(img *sane.Image, fileName string) error {
	tiffFile, err := os.Create(fileName + ".tiff")
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer tiffFile.Close()

	err = tiff.Encode(tiffFile, img, nil)
	if err != nil {
		return fmt.Errorf("error encoding image: %w", err)
	}
	return nil

}

func toPNG(img *sane.Image, fileName string) error {
	pngFile, err := os.Create(fileName + ".png")
	if err != nil {
		return err
	}
	defer pngFile.Close()

	err = png.Encode(pngFile, img)
	if err != nil {
		return err
	}
	return nil
}

func toTXT(img *sane.Image, fileName string) error {
	text, err := OCR(img)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName+"_ocr.txt", []byte(text), 0644)
}
