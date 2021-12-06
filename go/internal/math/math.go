package math

func Min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
func Max(a int, b int) int {
	if a < b {
		return b
	}

	return a
}
func Extremum(a int, b int) (int, int) {
	return Min(a, b), Max(a, b)
}
func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
