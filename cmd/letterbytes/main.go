// main package
package main

import (
	"fmt"
	"os"

	"github.com/waltervargas/letterbytes"
)

func validateScannerName() {
	scannerName := os.Getenv("LETTERBYTES_SCANNER")
	if scannerName == "" {
		fmt.Printf("❌ environment variable LETTERBYTES_SCANNER not set\n")
		fmt.Printf("🚀 finding scanners...\n")
		devices, err := letterbytes.FindScanner()
		if err != nil {
			panic(err)
		}
		fmt.Printf("possible values for env variable LETTERBYTES_SCANNER: \n\n")
		for _, d := range devices {
			fmt.Printf("LETTERBYTES_SCANNER='%s'\n", d.Name)
		}
		os.Exit(1)
	}
}

func main() {
	validateScannerName()
	err := letterbytes.Run(os.Getenv("LETTERBYTES_SCANNER"), os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Image scanned and saved to %s.png\n", os.Args[1])
	fmt.Printf("Image scanned and saved to %s.tiff\n", os.Args[1])
}
