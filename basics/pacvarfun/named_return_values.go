package pacvarfun

// sum int is a name of returning value
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
