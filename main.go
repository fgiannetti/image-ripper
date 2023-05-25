package main

import (
	"log"
	"os"

	"github.com/fgiannetti/image-ripper/pkg/finder"
	"github.com/fgiannetti/image-ripper/pkg/tar"
)

const tmp = "/tmp/image-ripper"
const filename = "/home/fer/Downloads/testTar/root.tar"
const dest = "/home/fer/Downloads/testTar/dest"

func main() {
	defer clean()

	//TODO: if the destination directory does not exists, create it. If it exists and is not empty, panic

	if err := os.Mkdir(tmp, 0777); os.IsExist(err) {
		log.Println("temp directory already exists. Deleting it...")
		if err = os.RemoveAll(tmp); err != nil {
			panic(err)
		}
	}

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	log.Printf("expanding the image tar file in the temporary directory %s", tmp)
	err = tar.Untar(file, tmp)

	if err != nil {
		panic(err)
	}

	log.Println("tar file expanded successfully")

	log.Println("finding layers...")
	layers, err := finder.FindLayers(tmp)
	if err != nil {
		panic(err)
	}

	log.Println(layers)
}

func clean() {
	if err := os.RemoveAll(tmp); err == nil {
		log.Println("temporary directory successfully removed")
	}
}
