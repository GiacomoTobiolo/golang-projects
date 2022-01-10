package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// make ascii map
	asciiMap := make(map[int][]string)
	// set rune count to space, and ignore first line
	c := 32
	firstLine := false
	// temporary var for ascii character
	var tempAscii []string

	bannerslice := []string{"standard", "shadow", "thinkertoy"}
	bannerToUse := ""

	if len(os.Args) != 4 {

		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\n\nEX: go run . something standard --output=<fileName.txt>")
		return
	}

	for _, banner := range bannerslice {
		if os.Args[2] == banner {
			bannerToUse = banner + ".txt"
		}
	}

	fileName := strings.Split(os.Args[3], "=")

	// error checking
	if fileName[0] != "--output" {

		// print error
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\n\nEX: go run . something standard --output=<fileName.txt>")

		// stop program
		return
	}

	file, err := os.Open(bannerToUse)
	if err != nil {
		log.Fatal(err)
	}
	// create output file
	destination, err := os.Create(fileName[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer destination.Close()

	// read file
	readFile := bufio.NewScanner(io.Reader(file))
	for readFile.Scan() {
		if readFile.Text() == "" && firstLine {
			// assign key to ascii val
			asciiMap[c] = tempAscii
			tempAscii = nil
			c++
		} else if firstLine {
			tempAscii = append(tempAscii, readFile.Text())
		}
		firstLine = true
	}
	asciiMap[c] = tempAscii

	// take the input from terminal and then split at occurance of newline (\n)
	input := strings.Split(os.Args[1], "\\n")

	// range through each word and print each character for each line
	var output []string

	for _, word := range input {
		for i := 0; i < 8; i++ {
			for j := 0; j < len(word); j++ {
				output = append(output, (asciiMap[int(word[j])][i]))
			}
			output = append(output, "\n")
		}
	}
	output = append(output, "\n")
	outputFile := bufio.NewWriter(destination)
	joinedfile := strings.Join(output, "")
	outputFile.WriteString(joinedfile)
	outputFile.Flush()
}
