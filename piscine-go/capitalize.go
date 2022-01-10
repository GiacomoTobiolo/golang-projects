package piscine

func Capitalize(s string) string {
	// sliceOfS := []rune(s)
	var lowerCaseSlice []rune
	l := len(s)
	for i := 0; i < l; i++ {
		lowerCaseSlice = append(lowerCaseSlice, 'a')
	}
	// Make all lowercase
	for index, char := range s {
		if char >= 'A' && char <= 'Z' {
			charString := ToLower(string(char))
			aSliceOfCharString := []rune(charString)
			lowerCaseSlice[index] = aSliceOfCharString[0]
		} else {
			lowerCaseSlice[index] = char
		}
	}
	if lowerCaseSlice[0] >= 'a' && lowerCaseSlice[0] <= 'z' {
		charString := ToUpper(string(lowerCaseSlice[0]))
		aSliceOfCharString := []rune(charString)
		lowerCaseSlice[0] = aSliceOfCharString[0]
	}
	for i := 1; i < len(lowerCaseSlice); i++ {
		if (lowerCaseSlice[i-1] < '0' || lowerCaseSlice[i-1] > '9') &&
			(lowerCaseSlice[i-1] < 'A' || lowerCaseSlice[i-1] > 'Z') &&
			(lowerCaseSlice[i-1] < 'a' || lowerCaseSlice[i-1] > 'z') {
			charString := ToUpper(string(lowerCaseSlice[i]))
			aSliceOfCharString := []rune(charString)
			lowerCaseSlice[i] = aSliceOfCharString[0]
		}
	}
	return string(lowerCaseSlice[0:l])
}
