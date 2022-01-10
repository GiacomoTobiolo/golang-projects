package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	colours := make(map[string]string)

	colours["reset"] = "\033[0m"
	colours["--color=red"] = "\033[31m"
	colours["--color=blue"] = "\033[34m"
	colours["--color=purple"] = "\033[35m"
	colours["--color=green"] = "\033[32m"
	colours["--color=yellow"] = "\033[33m"
	colours["--color=cyan"] = "\033[36m"
	colours["--color=white"] = "\033[37m"
	colours["--color=black"] = "\033[0;30m"
	colours["--color=brown"] = "\033[0;33m"
	colours["--color=grey"] = "\033[1;30m"
	colours["--color=orange"] = "\033[38;5;172m"
	colours["--color=pink"] = "\033[38;5;205m"

	/*for i := 0; i < len(colours); i++ {
	if colourChoice == colours {
		textColour == append(textColour, ColourMap[i])
	/*/

	// make ascii map
	asciiMap := make(map[int][]string)
	// set rune count to space, and ignore first line
	c := 32
	firstLine := false
	// temporary var for ascii character
	var tempAscii []string

	if len(os.Args) > 5 || len(os.Args) < 3 {

		fmt.Println("Usage: go run . [STRING] [OPTION]\n\nEX: go run . something --colour=<colour>")
		return
	}

	colourChoice := os.Args[2]

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

	filename := strings.Split(colourChoice, "=")

	if filename[0] != "--color" {
		fmt.Println("Usage: go run . [STRING] [OPTION]\n\nEX: go run . something --color=<color>")
		return
	}

	// take the input from terminal and then split at occurance of newline (\n)
	//input := strings.Split(os.Args[1], "\\n")
	input := strings.Split(os.Args[1], "\\n")

	charCol := -1
	charCol2 := -1

	if len(os.Args) == 4 {

		charCol, err = strconv.Atoi(os.Args[3])

		if err != nil {
			fmt.Println("Usage: go run . [STRING] [OPTION] [FIRSTELEMENT]")
		}

	}

	if len(os.Args) == 5 {
		charCol, err = strconv.Atoi(os.Args[3])
		charCol2, err = strconv.Atoi(os.Args[4])

		if err != nil {
			fmt.Println("Usage: go run . [STRING] [OPTION] [FIRSTELEMENT] [SECONDELEMENT")
		}
	}

	// range through each word and print each character for each line
	for _, word := range input {
		for i := 0; i < 8; i++ {
			for j := 0; j < len(word); j++ {
				if charCol == -1 && charCol2 == -1 {
					fmt.Print((colours[colourChoice]), asciiMap[int(word[j])][i])
					// fmt.Print((colours["reset"]), asciiMap[int(word[j])][i+1])
				} else if charCol != -1 && charCol2 == -1 {
					if j == charCol {
						fmt.Print((colours[colourChoice]), asciiMap[int(word[j])][i], colours["reset"])
					} else {
						fmt.Print(asciiMap[int(word[j])][i])
					}
				} else if j >= charCol && j < charCol2 {
					fmt.Print((colours[colourChoice]), asciiMap[int(word[j])][i], colours["reset"])

				} else {
					fmt.Print(asciiMap[int(word[j])][i])
				}

			}
			fmt.Println()
		}
	}

}
