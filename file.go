package main

import (
	"os"
)

func makeFilePath(file_name string) (string, error) {
	workfolder, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path_to_file := workfolder + "/" + file_name
	return path_to_file, nil
}
