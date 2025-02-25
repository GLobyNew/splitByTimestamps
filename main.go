package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type timestamp struct {
	name  string
	start string
	end   string
}

func printTimeStamps(t []timestamp) {
	for _, entry := range t {
		fmt.Printf("start time: %v, end time: %v, name: %v\n", entry.start, entry.end, entry.name)
	}
}

func readTimeStamps(file_name string) ([]timestamp, error) {
	timestamps := []timestamp{}

	path_to_file, err := makeFilePath(file_name)
	if err != nil {
		return []timestamp{}, err
	}
	file, err := os.Open(path_to_file)
	if err != nil {
		return []timestamp{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	firstScan := true

	for scanner.Scan() {
		line := scanner.Text()
		startTime, name, found := strings.Cut(line, " ")
		if !found {
			continue
		}
		timestamps = append(timestamps, timestamp{
			name:  name,
			start: startTime,
		})
		if !firstScan {
			timestamps[len(timestamps)-1].end = startTime
		}
		firstScan = false
	}

	timestamps[len(timestamps)-1].end, err = getLenFile(os.Args[1])

	if err != nil {
		return []timestamp{}, err
	}

	return timestamps, nil
}

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

}
