package piscine

func AppendRange(min, max int) []int {
	var myArray []int

	for i := 0; i < max-min; i++ {
		myArray = append(myArray, min+i)
	}
	return myArray
}
