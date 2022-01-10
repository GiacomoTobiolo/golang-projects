package piscine

// import "fmt"

func MakeRange(min, max int) []int {
	if max-min > 0 {
		myArray := make([]int, max-min)

		for i := 0; i < max-min; i++ {
			myArray[i] = min + i
		}
		return myArray
	}
	var nilSlice []int
	return nilSlice
}
