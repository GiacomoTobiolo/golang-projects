package piscine

func Rot14(s string) string {
	var sliceofRot14 []rune
	for _, char := range s {
		if char >= 'A' && char <= 'L' || char >= 'a' && char <= 'l' {
			sliceofRot14 = append(sliceofRot14, char+14)
		}
		if char > 'L' && char <= 'Z' {
			sliceofRot14 = append(sliceofRot14, char-'L'+'A'-1)
		}
		if char > 'l' && char <= 'z' {
			sliceofRot14 = append(sliceofRot14, char-'l'+'a'-1)
		}
		if char == ' ' || char == '?' || char >= '0' && char <= '9' {
			sliceofRot14 = append(sliceofRot14, char)
		}
	}
	return string(sliceofRot14)
}
