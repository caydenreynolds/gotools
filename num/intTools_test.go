package num

import "testing"

func TestMin(t *testing.T) {
	a := 10

	for b := 0; b < 10; b++ {
		if Min(a, b) != b {
			t.Errorf("Min function has failed")
		}
	}

	for b := 11; b < 20; b++ {
		if Min(a, b) != a {
			t.Errorf("Min function has failed")
		}
	}
}

func TestMax(t *testing.T) {
	a := 10

	for b := 0; b < 10; b++ {
		if Max(a, b) != a {
			t.Errorf("Max function has failed")
		}
	}

	for b := 11; b < 20; b++ {
		if Max(a, b) != b {
			t.Errorf("Max function has failed")
		}
	}
}
