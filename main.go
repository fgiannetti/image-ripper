package main

import (
	"os"
	"github.com/fgiannetti/image-ripper/pkg/tar"
)

const tmp = "/tmp/ripper"
const dest = "/tmp/dest"
const filename = "/Users/fernando.giannetti/Downloads/testTar/test.tar"

func main() {
	defer os.Remove(tmp)

	if err := os.Mkdir(tmp, 0777); err != nil {
		panic(err)
	}

	if err := os.Mkdir(tmp, 0777); err != nil {
		panic(err)
	}

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	//untar main image to temp
	err = tar.Untar(file, tmp)

	if err != nil {
		panic(err)
	}
}
