package service

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// Search files in directory by extension
func SearchFilesInDir(file_ext, path string) ([]string, error) {
	var files_list []string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) == file_ext {
			files_list = append(files_list, f.Name())
		}
	}

	if len(files_list) <= 0 {
		return nil, errors.New("No files found.")
	}

	return files_list, nil
}
