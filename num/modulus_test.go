package num

import (
	"testing"
)

func TestIntMod(t *testing.T) {
	if IntMod(2, 5) != 2 {
		t.Errorf("Expected 2 mod 5 to be 2. Instead got %v", IntMod(2, 5))
	}

	if IntMod(7, 5) != 2 {
		t.Errorf("Expected 7 mod 5 to be 2. Instead got %v", IntMod(7, 5))
	}

	if IntMod(5, 5) != 0 {
		t.Errorf("Expected 5 mod 5 to be 0. Instead got %v", IntMod(5, 5))
	}

	if IntMod(-22, -5) != 2 {
		t.Errorf("Expected -22 mod -5 to be 2. Instead got %v", IntMod(-22, -5))
	}
}
