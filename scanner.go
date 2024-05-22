package letterbytes

import (
	"fmt"

	"github.com/tjgq/sane"
)

func FindScanner() ([]sane.Device, error) {
	// Find the scanner
	err := sane.Init()
	if err != nil {
		return nil, err
	}
	return sane.Devices()
}

func Scan(scanner string, fileName string) error {
	img, err := scan(scanner)
	if err != nil {
		return err
	}

	err = toPNG(img, fileName)
	if err != nil {
		return err
	}

	err = toTIFF(img, fileName)
	if err != nil {
		return err
	}

	return nil
}

// name is string such as: escl:https://192.168.0.133:443
func scan(name string) (*sane.Image, error) {
	err := sane.Init()
	if err != nil {
		return nil, fmt.Errorf("error initializing SANE: %w", err)
	}
	defer sane.Exit()

	scanner, err := sane.Open(name)
	if err != nil {
		return nil, fmt.Errorf("error opening device: %w", err)
	}
	defer scanner.Close()

	_, err = scanner.SetOption("mode", "Color")
	if err != nil {
		return nil, err
	}
	_, err = scanner.SetOption("resolution", 150)
	if err != nil {
		return nil, err
	}

	// Start the scan
	img, err := scanner.ReadImage()
	if err != nil {
		return nil, fmt.Errorf("error during scan: %w", err)
	}

	return img, nil
}
