package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func removeRune(slice []rune, s int) []rune {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	//open file, read and split into an array of strings
	var inputSlice []string
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	//close file once it's finished
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		inputSlice = append(inputSlice, scanner.Text())
	}
	fmt.Println(inputSlice)
	for i := 0; i < len(inputSlice); i++ {
		if inputSlice[i] == "(up)" {
			inputSlice[i-1] = strings.ToUpper(inputSlice[i-1])
			inputSlice = remove(inputSlice, i)
			i--
		}
		if inputSlice[i] == "(low)" {
			inputSlice[i-1] = strings.ToLower(inputSlice[i-1])
			inputSlice = remove(inputSlice, i)
			i--
		}
		if inputSlice[i] == "(cap)" {
			inputSlice[i-1] = strings.Title(inputSlice[i-1])
			inputSlice = remove(inputSlice, i)
			i--
		}
		if inputSlice[i] == "(cap," {
			for j := 1; j <= int(inputSlice[i+1][0])-48; j++ {
				inputSlice[i-j] = strings.Title(inputSlice[i-j])
			}
			inputSlice = remove(inputSlice, i+1)
			inputSlice = remove(inputSlice, i)
			i--
			i--
		}
		if inputSlice[i] == "(up," {
			for j := 1; j <= int(inputSlice[i+1][0])-48; j++ {
				inputSlice[i-j] = strings.ToUpper(inputSlice[i-j])
			}
			inputSlice = remove(inputSlice, i+1)
			inputSlice = remove(inputSlice, i)
			i--
			i--
		}
		if inputSlice[i] == "(low," {
			for j := 1; j <= int(inputSlice[i+1][0])-48; j++ {
				inputSlice[i-j] = strings.ToLower(inputSlice[i-j])
			}
			inputSlice = remove(inputSlice, i+1)
			inputSlice = remove(inputSlice, i)
			i--
			i--
		}
		if inputSlice[i] == "(bin)" {
			value, _ := strconv.ParseInt(inputSlice[i-1], 2, 64)
			inputSlice[i-1] = fmt.Sprint(value)
			inputSlice = remove(inputSlice, i)
			i--
		}
		if inputSlice[i] == "(hex)" {
			value, _ := strconv.ParseInt(inputSlice[i-1], 16, 64)
			inputSlice[i-1] = fmt.Sprint(value)
			inputSlice = remove(inputSlice, i)
			i--
		}
	}
	for i := 0; i < len(inputSlice); i++ {
		if inputSlice[i] == "." || inputSlice[i] == "," || inputSlice[i] == "!" || inputSlice[i] == "?" || inputSlice[i] == ":" || inputSlice[i] == ";" {
			inputSlice[i-1] = inputSlice[i-1] + inputSlice[i]
			inputSlice = remove(inputSlice, i)
			i--
		}
		count := 0
		for p, srune := range inputSlice[i] {
			if srune == '.' || srune == ',' || srune == '!' || srune == '?' || srune == ':' || srune == ';' && p == 0 {
				count = 0
				for _, multiPunc := range inputSlice[i] {
					if multiPunc == '.' || multiPunc == ',' || multiPunc == '!' || multiPunc == '?' || multiPunc == ':' || multiPunc == ';' {
						count++
					}
				}
				if count == 1 && p == 0 {
					inputSlice[i-1] = inputSlice[i-1] + string(srune)
					strRune := []rune(inputSlice[i])
					strRune = removeRune(strRune, p)
					inputSlice[i] = string(strRune)
				} else if count == len(inputSlice[i]) {
					inputSlice[i-1] = inputSlice[i-1] + inputSlice[i]
					inputSlice = remove(inputSlice, i)
					i--
				}
			}
		}
	}
	found := false
	for i := 0; i < len(inputSlice); i++ {
		if inputSlice[i] == "'" {
			if !found {
				inputSlice[i+1] = inputSlice[i] + inputSlice[i+1]
				inputSlice = remove(inputSlice, i)
				i--
				found = true
			} else {
				inputSlice[i-1] = inputSlice[i-1] + inputSlice[i]
				inputSlice = remove(inputSlice, i)
				i--
				found = false
			}
		}
	}
	for i := 0; i < len(inputSlice); i++ {
		if inputSlice[i] == "a" || inputSlice[i] == "A" {
			if inputSlice[i+1][0] == 'a' || inputSlice[i+1][0] == 'e' || inputSlice[i+1][0] == 'i' || inputSlice[i+1][0] == 'o' || inputSlice[i+1][0] == 'u' || inputSlice[i+1][0] == 'h' {
				inputSlice[i] += "n"
			}
		}
	}
	file, err = os.Create(os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	outputString := strings.Join(inputSlice, " ")
	writer.WriteString(outputString)
	writer.Flush()
}
