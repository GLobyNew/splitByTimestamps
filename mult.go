package main

import (
	"errors"
	"os/exec"
	"strings"
)

func getLenFile(file_name string) (string, error) {
	file_path, err := makeFilePath(file_name)
	if err != nil {
		return "", err
	}

	app := "ffprobe"
	appArgs := []string{
		"-v", "error",
		"-show_entries", "format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1",
		"-sexagesimal",
		file_path,
	}
	cmd := exec.Command(app, appArgs...)

	result, err := cmd.Output()

	if err != nil {
		return "", err
	}

	trimmedRes, _, found := strings.Cut(string(result), ".")
	if !found {
		return "", errors.New("can't find end of file in mult.go")
	}

	return trimmedRes, nil

}
