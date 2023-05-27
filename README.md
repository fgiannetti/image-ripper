# Image Ripper

Exports a Docker image internal FileSystem to a local directory.

The tool recursivelly iterates through the directories of the root tar file looking for _layer.tar_ files and unpacking them inside the destination directory.

The result is a directory that contains all files and directories contained in the exported image.
The image should be a tar file

## Usage

### Build

```
make build
```

### Execute

```
./image-ripper <imageTarFile> <destinationDirectory> 
```


