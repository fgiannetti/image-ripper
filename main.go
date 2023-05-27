package main

import (
	"log"
	"os"

	"github.com/fgiannetti/image-ripper/pkg/ripper"
)

var returnCode = 0

func main() {
	ripper := ripper.New()
	defer ripper.Clean()

	if len(os.Args) < 3 {
		log.Println("usage: image-ripper <imageTarFile> <destinationDirectory>")
		returnCode = 1
		return
	}

	tarFile := os.Args[1]
	destDir := os.Args[2]

	if err := ripper.Unpack(tarFile, destDir); err != nil {
		returnCode = 1
		log.Printf("an error ocurred extracting image filesystem: %s", err.Error())
	}
}

func exit() {
	os.Exit(returnCode)
}
