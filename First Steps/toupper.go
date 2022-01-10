package piscine

func ToUpper(s string) string {
	aString := ""
	for _, i := range s {
		if i >= 97 && i <= 122 {
			i = i - 32
			aString += string(i)
		} else {
			aString += string(i)
		}
	}
	return aString
}
