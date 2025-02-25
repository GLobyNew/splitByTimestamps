package main

import (
	"os"
	"path"
)

func makeFilePath(file_name string) (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	workfolder := path.Dir(execPath)
	path_to_file := workfolder + "/" + file_name
	return path_to_file, nil
}
