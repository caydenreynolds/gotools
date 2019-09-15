package num

import (
	"gotools/src/num"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	mat := num.Matrix{}.NewMatrix(5, 4)

	m, n := mat.Dimensions()

	if m != 5 || n != 4 {
		t.Errorf("Expected matrix to have dimensions (5, 4). Instead got (%v, %v)", m, n)
	}

	if mat.GetValue(0, 0) != 0 {
		t.Errorf("Expected matrix (0, 0) to be 0. Instead got %v", mat.GetValue(0, 0))
	}

	if mat.GetValue(3, 2) != 0 {
		t.Errorf("Expected matrix (3, 2) to be 0. Instead got %v", mat.GetValue(3, 2))
	}

	if mat.GetValue(4, 3) != 0 {
		t.Errorf("Expected matrix (4, 3) to be 0. Instead got %v", mat.GetValue(4, 4))
	}
}

func TestNewMatrixFromSlice(t *testing.T) {
	slice := [][]float64{
		{1, 2, 3, 4, 5},
		{0.5, 1.5, 2.5, 3.5, 4.5},
		{0.9999, 7.435, 23.456, 3.256, 56.765},
	}
	mat := num.Matrix{}.NewMatrixFromSlice(slice)
	m, n := mat.Dimensions()

	if m != 3 || n != 5 {
		t.Errorf("Expected matrix to have dimensions (3, 5). Instead got (%v, %v)", m, n)
	}

	if !num.Approx(mat.GetValue(0, 0), 1) {
		t.Errorf("Expected matrix (0, 0) to be 1. Instead got %v", mat.GetValue(0, 0))
	}

	if !num.Approx(mat.GetValue(1, 2), 2.5) {
		t.Errorf("Expected matrix (1, 2) to be 2.5. Instead got %v", mat.GetValue(1, 2))
	}

	if !num.Approx(mat.GetValue(2, 4), 56.765) {
		t.Errorf("Expected matrix (2, 4) to be 56.765. Instead got %v", mat.GetValue(2, 4))
	}
}

func TestMatrixSlice(t *testing.T) {
	slice := [][]float64{
		{1, 2, 3, 4, 5},
		{0.5, 1.5, 2.5, 3.5, 4.5},
		{0.9999, 7.435, 23.456, 3.256, 56.765},
	}
	mat := num.Matrix{}.NewMatrixFromSlice(slice).Slice(1, 2, 0, 3)
	m, n := mat.Dimensions()

	if m != 1 || n != 3 {
		t.Errorf("Expected matrix to have dimensions (1, 3). Instead got (%v, %v)", m, n)
	}

	if !num.Approx(mat.GetValue(0, 0), 0.5) {
		t.Errorf("Expected matrix (0, 0) to be 0.5. Instead got %v", mat.GetValue(0, 0))
	}

	if !num.Approx(mat.GetValue(0, 2), 2.5) {
		t.Errorf("Expected matrix (0, 2) to be 2.5. Instead got %v", mat.GetValue(0, 2))
	}
}

func TestRotate(t *testing.T) {
	mat := num.Matrix{}.NewMatrix(1, 5)
	mat = mat.Rotate(true)
	m, n := mat.Dimensions()

	if m != 5 || n != 1 {
		t.Errorf("Expected matrix to have dimensions (5, 1). Instead got (%v, %v)", m, n)
	}

	mat = num.Matrix{}.NewMatrix(3, 1)
	mat = mat.Rotate(false)
	m, n = mat.Dimensions()

	if m != 1 || n != 3 {
		t.Errorf("Expected matrix to have dimensions (1, 3). Instead got (%v, %v)", m, n)
	}

	mat = num.Matrix{}.NewMatrix(2, 3)
	mat = mat.Rotate(true)
	m, n = mat.Dimensions()

	if m != 3 || n != 2 {
		t.Errorf("Expected matrix to have dimensions (3, 2). Instead got (%v, %v)", m, n)
	}

	slice := [][]float64{
		{1, 2},
		{3, 4},
	}
	mat = num.Matrix{}.NewMatrixFromSlice(slice)
	mat = mat.Rotate(false)
	m, n = mat.Dimensions()

	if m != 2 || n != 2 {
		t.Errorf("Expected matrix to have dimensions (3, 5). Instead got (%v, %v)", m, n)
	}

	if !num.Approx(mat.GetValue(0, 0), 2) {
		t.Errorf("Expected matrix (0, 0) to be 2. Instead got %v", mat.GetValue(0, 0))
	}

	if !num.Approx(mat.GetValue(0, 1), 4) {
		t.Errorf("Expected matrix (0, 1) to be 4. Instead got %v", mat.GetValue(0, 1))
	}

	if !num.Approx(mat.GetValue(1, 0), 1) {
		t.Errorf("Expected matrix (1, 0) to be 1. Instead got %v", mat.GetValue(1, 0))
	}

	if !num.Approx(mat.GetValue(1, 1), 3) {
		t.Errorf("Expected matrix (1, 1) to be 3. Instead got %v", mat.GetValue(1, 1))
	}
}

func TestDotProduct(t *testing.T) {
	slice1 := [][]float64{
		{1},
		{2},
		{3},
	}
	slice2 := [][]float64{
		{4, 5, 6},
	}

	mat1 := num.Matrix{}.NewMatrixFromSlice(slice1)
	mat2 := num.Matrix{}.NewMatrixFromSlice(slice2)

	if !num.Approx(mat1.DotProduct(mat2), 32) {
		t.Errorf("Expected dot product to be 32. Instead got %v", mat1.DotProduct(mat2))
	}

	slice1 = [][]float64{
		{1, 2, 3},
	}
	slice2 = [][]float64{
		{4},
		{5},
		{6},
	}

	mat1 = num.Matrix{}.NewMatrixFromSlice(slice1)
	mat2 = num.Matrix{}.NewMatrixFromSlice(slice2)

	if !num.Approx(mat1.DotProduct(mat2), 32) {
		t.Errorf("Expected dot product to be 32. Instead got %v", mat1.DotProduct(mat2))
	}
}

func TestMatrixMultiply(t *testing.T) {
	//Create matrices
	slice1 := [][]float64{
		{1, 2},
		{3, 4},
	}
	slice2 := [][]float64{
		{5, 6},
		{7, 8},
	}
	mat1 := num.Matrix{}.NewMatrixFromSlice(slice1)
	mat2 := num.Matrix{}.NewMatrixFromSlice(slice2)
	prod := mat1.Multiply(mat2)

	if m, n := prod.Dimensions(); m != 2 || n != 2 {
		t.Errorf("Expected result matrix to have dimensions (2, 2). Instead got (%v, %v)", m, n)
	}

	if !num.Approx(prod.GetValue(0, 0), 19) {
		t.Errorf("Expected result matrix to have value 19. Instead got %v", prod.GetValue(0, 0))
	}

	if !num.Approx(prod.GetValue(0, 1), 22) {
		t.Errorf("Expected result matrix to have value 22. Instead got %v", prod.GetValue(0, 1))
	}

	if !num.Approx(prod.GetValue(1, 0), 43) {
		t.Errorf("Expected result matrix to have value 43. Instead got %v", prod.GetValue(1, 0))
	}

	if !num.Approx(prod.GetValue(1, 1), 50) {
		t.Errorf("Expected result matrix to have value 50. Instead got %v", prod.GetValue(1, 1))
	}
}
