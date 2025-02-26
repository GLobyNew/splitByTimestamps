package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	programName := os.Args[0]
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf("Usage:\n\n\t%v <multimedia_file> <timestamps file>", programName)
		return
	}

	timestamps, err := readTimeStamps(args[1])
	if err != nil {
		log.Fatalln(err)
	}

	printTimeStamps(timestamps)
	sourceFilePath, err := makeFilePath(args[0])
	if err != nil {
		log.Fatalln(err)
	}
	err = processFile(sourceFilePath, timestamps)
	if err != nil {
		log.Fatalln(err)
	}
}
