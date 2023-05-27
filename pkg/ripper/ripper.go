package ripper

import (
	"fmt"
	"log"
	"os"

	"github.com/fgiannetti/image-ripper/pkg/finder"
	"github.com/fgiannetti/image-ripper/pkg/tar"
)

const tmp = "/tmp/image-ripper"

type ripper struct {
}

func New() *ripper {
	return &ripper{}
}

func (r *ripper) Unpack(tarFile string, destDir string) error {
	if err := prepareFS(tarFile, destDir); err != nil {
		return err
	}

	log.Printf("expanding the image tar file in the temporary directory %s", tmp)
	if err := tar.Untar(tarFile, tmp); err != nil {
		return err
	}

	log.Println("tar file expanded successfully")

	log.Println("finding layers...")
	layers, err := finder.FindLayers(tmp)
	if err != nil {
		return err
	}

	log.Printf("%d layers found", len(layers))

	for _, layer := range layers {
		if err := tar.Untar(layer, destDir); err != nil {
			return err
		}
	}

	log.Println("all layers were unpacked!")
	return nil
}

func (r *ripper) Clean() {
	if err := os.RemoveAll(tmp); err == nil {
		log.Println("temporary directory successfully removed")
	}
}

func prepareFS(tarFile string, destDir string) error {
	dest, err := os.Stat(destDir)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("destination directory %s does not exits", destDir)
		}
		return err
	}

	if !dest.IsDir() {
		return fmt.Errorf("destination directory path is not a directory")
	}

	if err := os.Mkdir(tmp, 0777); os.IsExist(err) {
		log.Println("temp directory already exists. Deleting it...")
		if err = os.RemoveAll(tmp); err != nil {
			panic(err)
		}
	}

	return nil
}
