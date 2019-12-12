package num

import "math"

//Computes the modulus of two integers
//Note: unlike the default go % operator, this function gives the expcted value when a is a negative number
func IntMod(a, b int) int {
	return a - b*int(math.Floor(float64(a)/float64(b)))
}
