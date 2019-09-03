package num

import "math"

func FindMax(floats []float32) float32 {
	var maximum float64
	for i := 0; i < len(floats); i++ {
		if i == 0 {
			maximum = float64(floats[i])
		} else {
			maximum = math.Max(maximum, float64(floats[i]))
		}
	}

	return float32(maximum)
}
