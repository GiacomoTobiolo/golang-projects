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
	/*var banner []string

	if os.Args[2] == "standard.txt"{
		banner = append(banner, "standard.txt")
	}else if os.Args[2] == "shadow.txt"{
		banner = append(banner, "shadow.txt")
	}else if os.Args[2] == "thinkertoy.txt"{
		banner = append(banner,"thinkertoy.txt")
	}

	*/

	bannerslice := []string{"standard", "shadow", "thinkertoy"}
	bannerToUse := ""

	for _, banner := range bannerslice {
		if os.Args[2] == banner {
			bannerToUse = banner + ".txt"
		}
	}

	if bannerToUse == "" {
		// print erro
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")

		// stop program
		return
	}

	file, err := os.Open(bannerToUse)
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
