package num

import "math"

//Find value of largest float in a slice
func FindMax(floats []float32) float32 {
	return floats[FindMaxIndex(floats)]
}

//Finds index of largest float in a slice
func FindMaxIndex(floats []float32) int {
	if len(floats) == 0 {
		panic("Cannot find maximum index of empty slice")
	}

	maximumValue := float64(floats[0])
	maximumIndex := 0

	for i := 1; i < len(floats); i++ {
		if i == 0 {
			maximumValue = float64(floats[i])
		} else {
			maximumIndex = i
			maximumValue = math.Max(maximumValue, float64(floats[i]))
		}
	}

	return maximumIndex
}
