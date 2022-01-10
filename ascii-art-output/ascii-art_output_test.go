/* Package contains the functions required to compare the command line output
with the expected result stored in a file.

https://github.com/nik-don/01-founders/tree/main/ascii-art/output
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// **************************
// NAME_OF_PROGRAM
const programName = "main.go"

// **************************

// All additional test cases will be read from this file
var fileName string = "test-cases.txt"

const errMsg = `Usage: go run . [OPTION]

EX: go run . --reverse=<fileName>`

type testData struct {
	arg      string
	flag     string
	expected string
}

var testCases = []testData{}

// TestAsciiArtOutput will test the output of the main.go program to match the desired representation in Ascii Art characters
func TestAsciiArtOutput(t *testing.T) {
	// add additional testcases from the file into the testCases
	addTestCases(fileName)

	var file string

	for _, test := range testCases {
		// check if flag contains the output flag and extract file name
		if test.flag != "" {
			if strings.Contains(test.flag, "--output=") {
				file = strings.SplitAfter(test.flag, "--output=")[1]
			}
		}

		cmd := "go run " + programName + " " + escapedString(test.arg) + " " + test.flag + " | cat -e"

		output, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			t.Fatal(">>> TRY CHECKING the input argument in the", fileName, " file. For >>>> ", test.arg, " \n", err)
		}

		cmd = "cat -e " + file

		fileOutput, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			t.Fatal(">>> TRY CHECKING the input argument in the", fileName, " file. For >>>> ", test.arg, " \n", err)
		}

		fileOutputStr := string(fileOutput)
		if test.flag == "" {
			fileOutputStr = string(output)
		}

		if fileOutputStr != (test.expected) {
			t.Errorf("For argument '%v' \n\033[1;34mexpected \n%s\033[0m\n\033[1;31mbut got \n%s\033[0m", test.arg, test.expected, fileOutputStr)
			panic("err")
		} else {
			t.Logf("output for \033[38;5;105m\"%s\"\033[39;49m is \n\033[1;32m%s\033[0m", test.arg, fileOutputStr)
			t.Log("\n")
			t.Logf("expected output for \033[38;5;105m\"%s\"\033[39;49m is \n\033[1;32m%s\033[0m\n\033[1;36m------------------------------\033[0m\n", test.arg+test.flag, test.expected)
		}
	}
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

// addTestCases will read the contents of the specified file, extract and add the desired output into the testCases
func addTestCases(fileName string) {
	lines, err := readLines(fileName)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var (
		arg     string
		file    string
		tempStr string
	)

	for i, line := range lines {

		if len(line) > 0 && line[0] == '#' {

			if tempStr != "" {

				testCases = append(testCases, testData{arg, file, tempStr})

				tempStr = ""
			}
			arg, file = extract(string(line[1 : len(line)-1]))

		} else {
			tempStr += line + "\n"
		}
		// Last test case(No more testcases in the file)
		if i == len(lines)-1 {
			testCases = append(testCases, testData{arg, file, tempStr})
		}
	}
}

/*extract

//go run . "banana standard abc" shadow

a, b := extract(arg)

fmt.Println(a) // banana standard abc
fmt.Println(b) // shadow
*/
func extract(s string) (string, string) {
	a, b := "", ""

	// double quotes
	fi := strings.Index(s, "\"")
	li := strings.LastIndex(s, "\"")

	// catch slice bounds out of range errors
	if fi <= len(s) {
		a = s[fi+1 : li]
	}

	// rest of string
	if li+2 < len(s) {
		b = s[li+2:]
	}

	return a, b
}

// escapedString is used to pass a string that is already escaped as an argument into the command line
// example:	"1a\"#F"
// becomes:	"1a"#F"
func escapedString(s string) string {
	s = fmt.Sprint(s)
	return ("\"" + s + "\"")
}
