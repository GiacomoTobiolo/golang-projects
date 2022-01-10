package piscine

func ToLower(s string) string {
	aString := ""
	for _, i := range s {
		if i >= 65 && i <= 90 {
			i = i + 32
			aString += string(i)
		} else {
			aString += string(i)
		}
	}
	return aString
}
