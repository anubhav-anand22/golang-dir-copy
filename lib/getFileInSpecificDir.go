package lib

import (
	"io/fs"
	"os"
)

func GetItemsInFolder(folderPath string) ([]fs.FileInfo, error) {
	// Open the folder
	dir, err := os.Open(folderPath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	// Read the contents of the folder
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	return fileInfos, nil
}
