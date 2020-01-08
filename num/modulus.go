package num

import "math"

//Computes the modulus of two integers
//Note: unlike the default go % operator, this function gives the expected value when a is a negative number
func IntMod(a, b int) int {
	return int(math.Abs(float64(a - b*(a/b))))
}
