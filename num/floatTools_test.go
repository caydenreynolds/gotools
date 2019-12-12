package num

import (
	"testing"
)

func TestFindMaxIndex(t *testing.T) {
	floats := []float64{1, 5, 9, 7, 5}

	if FindMaxIndex(floats) != 2 {
		t.Errorf("Expected index 2, instead got %v", FindMaxIndex(floats))
	}
}

func TestFindMax(t *testing.T) {
	floats := []float64{1, 5, 9, 7, 4}

	if FindMax(floats) != 9 {
		t.Errorf("Expected output 9, instead got %v", FindMax(floats))
	}
}

func TestApprox(t *testing.T) {

	if !Approx(1.0, 2.0/2) {
		t.Errorf("TestApprox failed for 1, 2/2")
	}
	if !Approx(1.0/8, 0.125) {
		t.Errorf("TestApprox failed for 1/3, 0.33")
	}
	if !Approx(-1.0/4, -0.25) {
		t.Errorf("TestApprox failed for -1/4, -0.25")
	}
	if Approx(1.0/4, 0.4) {
		t.Errorf("TestApprox failed for 1/6, 0.33")
	}
}
