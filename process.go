package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/schollz/progressbar/v3"
)

const (
	outFolder = "splitOutput"
)

func processTrack(filePath string, ts timestamp, ch chan struct{}) {
	app := "ffmpeg"
	outFilePath := fmt.Sprintf("%v/%v.mp3", outFolder, ts.name)
	appArgs := []string{
		"-ss", ts.start,
		"-to", ts.end,
		"-i", filePath,
		outFilePath,
	}
	exec.Command(app, appArgs...)
	cmd := exec.Command(app, appArgs...)

	_, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}
	<-ch
}

func processFile(filePath string, timestamps []timestamp) error {
	err := os.Mkdir(outFolder, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	ch := make(chan struct{})
	for _, entry := range timestamps {
		go processTrack(filePath, entry, ch)
	}
	bar := progressbar.Default(int64(len(timestamps)))
	for range len(timestamps) {
		ch <- struct{}{}
		bar.Add(1)
	}

	return nil
}
