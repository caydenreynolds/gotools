package num

import (
	"gotools/src/num"
	"testing"
)

func TestFindMaxIndex(t *testing.T) {
	floats := []float64{1, 5, 9, 7, 5}

	if num.FindMaxIndex(floats) != 2 {
		t.Errorf("Expected index 2, instead got %v", num.FindMaxIndex(floats))
	}
}

func TestFindMax(t *testing.T) {
	floats := []float64{1, 5, 9, 7, 4}

	if num.FindMax(floats) != 9 {
		t.Errorf("Expected output 9, instead got %v", num.FindMax(floats))
	}
}

func TestApprox(t *testing.T) {

	if !num.Approx(1.0, 2.0/2) {
		t.Errorf("TestApprox failed for 1, 2/2")
	}
	if !num.Approx(1.0/8, 0.125) {
		t.Errorf("TestApprox failed for 1/3, 0.33")
	}
	if !num.Approx(-1.0/4, -0.25) {
		t.Errorf("TestApprox failed for -1/4, -0.25")
	}
	if num.Approx(1.0/4, 0.4) {
		t.Errorf("TestApprox failed for 1/6, 0.33")
	}
}
