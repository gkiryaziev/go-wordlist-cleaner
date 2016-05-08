package service

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// SearchFilesInDir return list of files in directory by extension
func SearchFilesInDir(fileExt, path string) ([]string, error) {
	var filesList []string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) == fileExt {
			filesList = append(filesList, f.Name())
		}
	}

	if len(filesList) <= 0 {
		return nil, errors.New("No files found.")
	}

	return filesList, nil
}
