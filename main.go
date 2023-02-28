package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("expected csv file to chunk-ify")
	}

	// open file
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// read the first line
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		// TODO errors.wrap or something
		log.Fatalln("unable to read header line")
		log.Fatal(err)
	}

	header := scanner.Text()

	linecount := 0
	currentFileIndex := 1
	currentFileName := fmt.Sprintf("%s-split-%d.csv", os.Args[1], currentFileIndex)
	currentFile, err := os.Create(currentFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer currentFile.Close()

	currentFile.WriteString(header)
	currentFile.WriteString("\n")

	// read the rest
	for scanner.Scan() {
		line := scanner.Bytes()
		linecount += 1
		currentFile.Write(line)
		currentFile.WriteString("\n")

		// open next file after 500 lines
		if linecount >= 1000 {
			linecount = 0
			currentFile.Close()

			currentFileIndex += 1
			currentFileName = fmt.Sprintf("%s-split-%d.csv", os.Args[1], currentFileIndex)
			currentFile, err = os.Create(currentFileName)

			if err != nil {
				log.Fatal(err)
			}
			defer currentFile.Close()

			currentFile.WriteString(header)
			currentFile.WriteString("\n")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
