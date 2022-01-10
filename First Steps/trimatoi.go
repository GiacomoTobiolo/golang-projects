package piscine

func TrimAtoi(s string) int {
	strrune := []rune(s)
	var ints []rune
	bool1 := false
	bool2 := false
	n := 0
	for y := 0; y < len(strrune); y++ {
		if strrune[y] == '-' && !bool2 {
			bool1 = true
		}
		if strrune[y] >= '0' && strrune[y] <= '9' {
			if !bool1 {
				bool2 = true
			}
		}

	}

	// fmt.Println(bool1)
	// fmt.Println(bool2)
	for j := 0; j < len(strrune); j++ {
		if strrune[j] >= '0' && strrune[j] <= '9' {
			ints = append(ints, strrune[j])
		}
	}
	for i := 0; i < len(ints); i++ {
		y := ints[i] - '0'
		n = n*10 + int(y)
	}
	if bool1 {
		n = -n
	}
	return n
}
