package piscine

// import "fmt"

func ConcatParams(args []string) string {
	var s string

	for i, str := range args {
		if i < len(args)-1 {
			s += str + "\n"
		} else {
			s += str
		}
	}
	return s
}
