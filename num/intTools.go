package num

import "math"

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func Abs(n int) int {
	return int(math.Abs(float64(n)))
}
