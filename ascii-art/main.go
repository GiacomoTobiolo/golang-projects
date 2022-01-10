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

	// open file
	file, err := os.Open("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
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

	for _, word := range input {
		for i := 0; i < 8; i++ {
			for j := 0; j < len(word); j++ {
				fmt.Print(asciiMap[int(word[j])][i])
			}
			fmt.Println()
		}
	}
}
