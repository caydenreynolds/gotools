package unique

import (
	"gotools/uniques/unique/src"
	"testing"
)

//Ensures first value of
func TestIncrStrFirstValue(t *testing.T) {
	uniqueStr := src.Str{}
	println(len(uniqueStr.Bytes))

	output := uniqueStr.Next()

	if output != "" {
		t.Errorf("First value of increment bytes incorrect. Expected , got " + output)
	}
}

func TestIncrStrFirst1000Values(t *testing.T) {
	valueCount := 1000
	uniqueStr := src.Str{}
	set := make(map[string]bool)

	for i := 0; i < valueCount; i++ {
		set[uniqueStr.Next()] = true
	}

	if len(set) != valueCount {
		t.Errorf("Not all of the first 1000 strings are unique")
	}
}

func TestIncrStrManyValues(t *testing.T) {
	valueCount := 1000000
	uniqueStr := src.Str{}
	set := make(map[string]bool)

	for i := 0; i < valueCount; i++ {
		set[uniqueStr.Next()] = true
	}

	if len(set) != valueCount {
		t.Errorf("Not all of the first 1000 strings are unique")
	}
}
