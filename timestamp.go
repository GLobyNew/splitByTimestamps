package main

import (
	"bufio"
	"errors"
	"fmt"
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

func formatStrTimestamp(t string) (string, error) {
	splTimestamps := strings.Split(t, ":")

	var sanitizeTime = func(s string) string {
		if len(s) < 2 {
			return "0" + s
		}
		return s
	}

	switch len(splTimestamps) {
	case 2:
		return fmt.Sprintf("00:%v:%v",
			sanitizeTime(splTimestamps[0]),
			sanitizeTime(splTimestamps[1])), nil
	case 3:
		return fmt.Sprintf("%v:%v:%v",
			sanitizeTime(splTimestamps[0]),
			sanitizeTime(splTimestamps[1]),
			sanitizeTime(splTimestamps[2])), nil
	default:
		return "", errors.New("can't format timestamp correctly")
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
		start, err := formatStrTimestamp(startTime)
		if !found {
			continue
		}
		if err != nil {
			return []timestamp{}, err
		}
		timestamps = append(timestamps, timestamp{
			name:  name,
			start: start,
		})
		if !firstScan {
			timestamps[len(timestamps)-2].end = start
		}
		firstScan = false
	}

	timestamps[len(timestamps)-1].end, err = getLenFile(os.Args[1])

	if err != nil {
		return []timestamp{}, err
	}

	return timestamps, nil
}
