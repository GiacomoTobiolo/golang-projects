package piscine

func Abort(a, b, c, d, e int) int {
	var intarr []int
	intarr = append(intarr, a, b, c, d, e)
	for index := range intarr {
		for index2 := range intarr {
			if intarr[index] <= intarr[index2] {
				intarr[index2], intarr[index] = intarr[index], intarr[index2]
			}
		}
	}
	return intarr[2]
}
