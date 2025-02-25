package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

type timestamp struct {
	name  string
	start string
	end   string
}

func printTimeStamps(t []timestamp) {
	for _, entry := range t {
		fmt.Printf("time: %v name: %v\n", entry.start, entry.name)
	}
}

func readTimeStamps(file_name string) ([]timestamp, error) {
	timestamps := []timestamp{}
	execPath, err := os.Executable()
	if err != nil {
		return []timestamp{}, err
	}
	workfolder := path.Dir(execPath)
	path_to_file := workfolder + "/" + file_name
	fmt.Println(path_to_file)

	file, err := os.Open(path_to_file)
	if err != nil {
		return []timestamp{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		time, name, found:= strings.Cut(line, " ")
		if !found {
			continue
		}
		timestamps = append(timestamps, timestamp{
			name:  name,
			start: time,
		})
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
