package num

import "math"

var epsilon = math.Nextafter(1.0, 2.0) - 1.0

//Find value of largest float in a slice
func FindMax(floats []float64) float64 {
	return floats[FindMaxIndex(floats)]
}

//Finds index of largest float in a slice
func FindMaxIndex(floats []float64) int {
	if len(floats) == 0 {
		panic("Cannot find maximum index of empty slice")
	}

	maximumValue := floats[0]
	maximumIndex := 0

	for i := 1; i < len(floats); i++ {
		if floats[i] > maximumValue {
			maximumValue = floats[i]
			maximumIndex = i
		}
	}

	return maximumIndex
}

//Returns true if two floats are approximately equal, returns false otherwise
func Approx(f1, f2 float64) bool {
	return math.Abs(f1-f2) <= epsilon
}
