package finder

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const layerFileName = "layer.tar"

func FindLayers(rootDir string) ([]string, error) {
	root, err := os.Stat(rootDir)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("directory %s do not exists for layer files search", root)
		}

		return nil, err
	}

	if !root.IsDir() {
		return nil, fmt.Errorf("the root directory to find must not be a file")
	}

	layerPaths := make([]string, 0)

	if err = filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if d.Name() == layerFileName {
			layerPaths = append(layerPaths, path)
			return filepath.SkipDir
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return layerPaths, nil
}
