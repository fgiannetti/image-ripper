package finder

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const testPath = "../../test/finderTest"
const subdir = testPath + "/subdir"

func TestFindOnRoot(t *testing.T) {
	setup(true)
	defer clean()

	os.Create(testPath + "/" + layerFileName)

	paths, err := FindLayers(testPath)

	require.NoError(t, err)
	require.NotEmpty(t, paths)
	require.Equal(t, 1, len(paths))
}

func TestFindOnSubDir(t *testing.T) {
	setup(true)
	defer clean()

	os.Create(subdir + "/" + layerFileName)

	paths, err := FindLayers(testPath)

	require.NoError(t, err)
	require.NotEmpty(t, paths)
	require.Equal(t, 1, len(paths))
}

func TestFindBothRootAndSubDir(t *testing.T) {
	setup(true)
	defer clean()

	os.Create(testPath + "/" + layerFileName)
	os.Create(subdir + "/" + layerFileName)

	paths, err := FindLayers(testPath)

	require.NoError(t, err)
	require.NotEmpty(t, paths)
	require.Equal(t, 1, len(paths))
}

func TestFindOnRootNotFound(t *testing.T) {
	setup(false)
	defer clean()

	paths, err := FindLayers(testPath)

	require.NoError(t, err)
	require.Empty(t, paths)
}

func setup(addLayerFile bool) {
	os.Mkdir(testPath, 0777)
	os.Mkdir(subdir, 0777)
	os.Create(testPath + "/file1.txt")
	os.Create(subdir + "/file2.txt")
}

func clean() {
	os.RemoveAll(testPath)
}
