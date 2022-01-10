package piscine

// import "fmt"

func SplitWhiteSpaces(s string) []string {
	var s0 string
	var myArray []string

	for i, char := range s {
		if char != ' ' && char != 9 && char != '\n' {
			s0 += string(char)
		} else if s[i-1] != ' ' && s[i-1] != 9 && s[i-1] != '\n' {
			myArray = append(myArray, s0)
			s0 = ""
		}
		if i == len(s)-1 {
			myArray = append(myArray, s0)
		}
	}
	return myArray
}
