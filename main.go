package main

import (
	"os"
	"github.com/fgiannetti/image-ripper/pkg/tar"
)

const tmp = "/tmp/image-ripper"
const filename = "/Users/fernando.giannetti/Downloads/testTar/test.tar"

func main() {
	defer clean()

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

func clean(){
	os.Remove(tmp)
}
