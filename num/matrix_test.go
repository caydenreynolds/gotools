package num

import (
	"testing"
)

func TestNewMatrix(t *testing.T) {
	mat := Matrix{}.NewMatrix(5, 4)

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
	mat := Matrix{}.NewMatrixFromSlice(slice)
	m, n := mat.Dimensions()

	if m != 3 || n != 5 {
		t.Errorf("Expected matrix to have dimensions (3, 5). Instead got (%v, %v)", m, n)
	}

	if !Approx(mat.GetValue(0, 0), 1) {
		t.Errorf("Expected matrix (0, 0) to be 1. Instead got %v", mat.GetValue(0, 0))
	}

	if !Approx(mat.GetValue(1, 2), 2.5) {
		t.Errorf("Expected matrix (1, 2) to be 2.5. Instead got %v", mat.GetValue(1, 2))
	}

	if !Approx(mat.GetValue(2, 4), 56.765) {
		t.Errorf("Expected matrix (2, 4) to be 56.765. Instead got %v", mat.GetValue(2, 4))
	}
}

func TestCopyConstructor(t *testing.T) {
	slice := [][]float64{
		{1},
	}
	mat1 := Matrix{}.NewMatrixFromSlice(slice)
	mat2 := mat1.Clone()
	mat2.SetValue(0, 0, 0)

	if !Approx(mat1.GetValue(0, 0), 1) {
		t.Errorf("Expected matrix1 (0, 0) to be 1. Instead got %v", mat1.GetValue(0, 0))
	}

	if !Approx(mat2.GetValue(0, 0), 0) {
		t.Errorf("Expected matrix (0, 0) to be 0. Instead got %v", mat2.GetValue(0, 0))
	}
}

func TestMatrixSlice(t *testing.T) {
	slice := [][]float64{
		{1, 2, 3, 4, 5},
		{0.5, 1.5, 2.5, 3.5, 4.5},
		{0.9999, 7.435, 23.456, 3.256, 56.765},
	}
	mat := Matrix{}.NewMatrixFromSlice(slice).Slice(1, 2, 0, 3)
	m, n := mat.Dimensions()

	if m != 1 || n != 3 {
		t.Errorf("Expected matrix to have dimensions (1, 3). Instead got (%v, %v)", m, n)
	}

	if !Approx(mat.GetValue(0, 0), 0.5) {
		t.Errorf("Expected matrix (0, 0) to be 0.5. Instead got %v", mat.GetValue(0, 0))
	}

	if !Approx(mat.GetValue(0, 2), 2.5) {
		t.Errorf("Expected matrix (0, 2) to be 2.5. Instead got %v", mat.GetValue(0, 2))
	}
}

func TestRotate(t *testing.T) {
	mat := Matrix{}.NewMatrix(1, 5)
	mat = mat.Rotate(true)
	m, n := mat.Dimensions()

	if m != 5 || n != 1 {
		t.Errorf("Expected matrix to have dimensions (5, 1). Instead got (%v, %v)", m, n)
	}

	mat = Matrix{}.NewMatrix(3, 1)
	mat = mat.Rotate(false)
	m, n = mat.Dimensions()

	if m != 1 || n != 3 {
		t.Errorf("Expected matrix to have dimensions (1, 3). Instead got (%v, %v)", m, n)
	}

	mat = Matrix{}.NewMatrix(2, 3)
	mat = mat.Rotate(true)
	m, n = mat.Dimensions()

	if m != 3 || n != 2 {
		t.Errorf("Expected matrix to have dimensions (3, 2). Instead got (%v, %v)", m, n)
	}

	slice := [][]float64{
		{1, 2},
		{3, 4},
	}
	mat = Matrix{}.NewMatrixFromSlice(slice)
	mat = mat.Rotate(false)
	m, n = mat.Dimensions()

	if m != 2 || n != 2 {
		t.Errorf("Expected matrix to have dimensions (3, 5). Instead got (%v, %v)", m, n)
	}

	if !Approx(mat.GetValue(0, 0), 2) {
		t.Errorf("Expected matrix (0, 0) to be 2. Instead got %v", mat.GetValue(0, 0))
	}

	if !Approx(mat.GetValue(0, 1), 4) {
		t.Errorf("Expected matrix (0, 1) to be 4. Instead got %v", mat.GetValue(0, 1))
	}

	if !Approx(mat.GetValue(1, 0), 1) {
		t.Errorf("Expected matrix (1, 0) to be 1. Instead got %v", mat.GetValue(1, 0))
	}

	if !Approx(mat.GetValue(1, 1), 3) {
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

	mat1 := Matrix{}.NewMatrixFromSlice(slice1)
	mat2 := Matrix{}.NewMatrixFromSlice(slice2)

	if !Approx(mat1.DotProduct(mat2), 32) {
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

	mat1 = Matrix{}.NewMatrixFromSlice(slice1)
	mat2 = Matrix{}.NewMatrixFromSlice(slice2)

	if !Approx(mat1.DotProduct(mat2), 32) {
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
	mat1 := Matrix{}.NewMatrixFromSlice(slice1)
	mat2 := Matrix{}.NewMatrixFromSlice(slice2)
	prod := mat1.Multiply(mat2)

	if m, n := prod.Dimensions(); m != 2 || n != 2 {
		t.Errorf("Expected result matrix to have dimensions (2, 2). Instead got (%v, %v)", m, n)
	}

	if !Approx(prod.GetValue(0, 0), 19) {
		t.Errorf("Expected result matrix to have value 19. Instead got %v", prod.GetValue(0, 0))
	}

	if !Approx(prod.GetValue(0, 1), 22) {
		t.Errorf("Expected result matrix to have value 22. Instead got %v", prod.GetValue(0, 1))
	}

	if !Approx(prod.GetValue(1, 0), 43) {
		t.Errorf("Expected result matrix to have value 43. Instead got %v", prod.GetValue(1, 0))
	}

	if !Approx(prod.GetValue(1, 1), 50) {
		t.Errorf("Expected result matrix to have value 50. Instead got %v", prod.GetValue(1, 1))
	}
}

func TestMatrixAdd(t *testing.T) {
	//Create matrices
	slice1 := [][]float64{
		{1, 2},
		{3, 4},
	}
	slice2 := [][]float64{
		{5, 6},
		{7, 8},
	}
	mat1 := Matrix{}.NewMatrixFromSlice(slice1)
	mat2 := Matrix{}.NewMatrixFromSlice(slice2)
	prod := mat1.Add(mat2)

	if m, n := prod.Dimensions(); m != 2 || n != 2 {
		t.Errorf("Expected result matrix to have dimensions (2, 2). Instead got (%v, %v)", m, n)
	}

	if !Approx(prod.GetValue(0, 0), 6) {
		t.Errorf("Expected result matrix to have value 6. Instead got %v", prod.GetValue(0, 0))
	}

	if !Approx(prod.GetValue(0, 1), 8) {
		t.Errorf("Expected result matrix to have value 8. Instead got %v", prod.GetValue(0, 1))
	}

	if !Approx(prod.GetValue(1, 0), 10) {
		t.Errorf("Expected result matrix to have value 10. Instead got %v", prod.GetValue(1, 0))
	}

	if !Approx(prod.GetValue(1, 1), 12) {
		t.Errorf("Expected result matrix to have value 12. Instead got %v", prod.GetValue(1, 1))
	}
}

func TestMatrix_Subtract(t *testing.T) {
	//Create matrices
	slice1 := [][]float64{
		{1, 2},
		{3, 4},
	}
	slice2 := [][]float64{
		{5, 6},
		{7, 8},
	}
	mat1 := Matrix{}.NewMatrixFromSlice(slice1)
	mat2 := Matrix{}.NewMatrixFromSlice(slice2)
	prod := mat1.Subtract(mat2)

	if m, n := prod.Dimensions(); m != 2 || n != 2 {
		t.Errorf("Expected result matrix to have dimensions (2, 2). Instead got (%v, %v)", m, n)
	}

	if !Approx(prod.GetValue(0, 0), -4) {
		t.Errorf("Expected result matrix to have value -4. Instead got %v", prod.GetValue(0, 0))
	}

	if !Approx(prod.GetValue(0, 1), -4) {
		t.Errorf("Expected result matrix to have value -4. Instead got %v", prod.GetValue(0, 1))
	}

	if !Approx(prod.GetValue(1, 0), -4) {
		t.Errorf("Expected result matrix to have value -4. Instead got %v", prod.GetValue(1, 0))
	}

	if !Approx(prod.GetValue(1, 1), -4) {
		t.Errorf("Expected result matrix to have value -4. Instead got %v", prod.GetValue(1, 1))
	}
}

func TestMatrix_Scale(t *testing.T) {
	//Create matrices
	slice1 := [][]float64{
		{1, 2},
		{3, 4},
	}
	mat1 := Matrix{}.NewMatrixFromSlice(slice1)
	prod := mat1.Scale(1 / 10.0)

	if m, n := prod.Dimensions(); m != 2 || n != 2 {
		t.Errorf("Expected result matrix to have dimensions (2, 2). Instead got (%v, %v)", m, n)
	}

	if !Approx(prod.GetValue(0, 0), 0.1) {
		t.Errorf("Expected result matrix to have value 0.1. Instead got %v", prod.GetValue(0, 0))
	}

	if !Approx(prod.GetValue(0, 1), 0.2) {
		t.Errorf("Expected result matrix to have value 0.2. Instead got %v", prod.GetValue(0, 1))
	}

	if !Approx(prod.GetValue(1, 0), 0.3) {
		t.Errorf("Expected result matrix to have value 0.3. Instead got %v", prod.GetValue(1, 0))
	}

	if !Approx(prod.GetValue(1, 1), 0.4) {
		t.Errorf("Expected result matrix to have value 0.4. Instead got %v", prod.GetValue(1, 1))
	}
}

func TestMatrix_Transpose(t *testing.T) {
	//Create matrix
	slice1 := [][]float64{
		{1, 2},
		{3, 4},
	}
	mat := Matrix{}.NewMatrixFromSlice(slice1)
	mat = mat.Transpose()

	if m, n := mat.Dimensions(); m != 2 || n != 2 {
		t.Errorf("Expected result matrix to have dimensions (2, 2). Instead got (%v, %v)", m, n)
	}

	if !Approx(mat.GetValue(0, 0), 1) {
		t.Errorf("Expected result matrix to have value 1. Instead got %v", mat.GetValue(0, 0))
	}

	if !Approx(mat.GetValue(0, 1), 3) {
		t.Errorf("Expected result matrix to have value 3. Instead got %v", mat.GetValue(0, 1))
	}

	if !Approx(mat.GetValue(1, 0), 2) {
		t.Errorf("Expected result matrix to have value 2. Instead got %v", mat.GetValue(1, 0))
	}

	if !Approx(mat.GetValue(1, 1), 4) {
		t.Errorf("Expected result matrix to have value 4. Instead got %v", mat.GetValue(1, 1))
	}
}

func TestMatrix_ExpandVector(t *testing.T) {
	slice := [][]float64{
		{1, 2},
	}
	vect := Matrix{}.NewMatrixFromSlice(slice)
	mat := vect.ExpandVector(2)

	if m, n := mat.Dimensions(); m != 2 || n != 2 {
		t.Errorf("Expected result matrix to have dimensions (2, 2). Instead got (%v, %v)", m, n)
	}

	if !Approx(mat.GetValue(0, 0), 1) {
		t.Errorf("Expected result matrix to have value 1. Instead got %v", mat.GetValue(0, 0))
	}

	if !Approx(mat.GetValue(0, 1), 2) {
		t.Errorf("Expected result matrix to have value 2. Instead got %v", mat.GetValue(0, 1))
	}

	if !Approx(mat.GetValue(1, 0), 1) {
		t.Errorf("Expected result matrix to have value 1. Instead got %v", mat.GetValue(1, 0))
	}

	if !Approx(mat.GetValue(1, 1), 2) {
		t.Errorf("Expected result matrix to have value 2. Instead got %v", mat.GetValue(1, 1))
	}
}
