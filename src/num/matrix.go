package num

import "fmt"

//Represents a matrix
type Matrix struct {
	values [][]float64
}

//Create am e,pty M x N matrix
func (mat Matrix) NewMatrix(m, n int) *Matrix {
	mat.values = make([][]float64, m)
	for i := range mat.values {
		mat.values[i] = make([]float64, n)
	}
	return &mat
}

//Create a matrix from a two dimensional slice. The second dimension of the slice must all be of the same length
func (mat Matrix) NewMatrixFromSlice(values [][]float64) *Matrix {
	//Validate inputs
	m := -1
	for _, value := range values {
		if m == -1 {
			m = len(value)
		} else if m != len(value) {
			panic(fmt.Sprintf("Cannot create matrices composed of vectors of different sizes. "+
				"Expected size %v, instead got size %v", m, len(value)))
		}
	}

	mat.values = make([][]float64, len(values))
	for i := range mat.values {
		mat.values[i] = make([]float64, len(values[0]))
		copy(mat.values[i], values[i])
	}

	return &mat
}

func (mat *Matrix) CopyMatrix() *Matrix {
	return Matrix{}.NewMatrixFromSlice(mat.values)
}

func (mat *Matrix) Slice(m1, m2, n1, n2 int) *Matrix {
	slices := mat.values[m1:m2]
	newValues := make([][]float64, len(slices))
	for i, value := range slices {
		newValues[i] = value[n1:n2]
	}
	return Matrix{}.NewMatrixFromSlice(newValues)
}

//Returns the value stored in the matrix at position (m, n)
func (mat *Matrix) GetValue(m, n int) float64 {
	return mat.values[m][n]
}

func (mat *Matrix) SetValue(m int, n int, value float64) {
	mat.values[m][n] = value
}

func (mat *Matrix) Dimensions() (int, int) {
	return len(mat.values), len(mat.values[0])
}

//Performs matrix multiplication
func (mat *Matrix) Multiply(other *Matrix) *Matrix {
	m, n := mat.Dimensions()
	o, p := other.Dimensions()
	if n != o {
		panic(fmt.Sprintf("Cannot compute multiply Matrices of different column/row sizes. "+
			"Base Matrix width = %v. Other Matrix height = %v", n, o))
	}

	result := Matrix{}.NewMatrix(m, p)

	for i := 0; i < m; i++ {
		for j := 0; j < p; j++ {
			slice1 := mat.Slice(i, i+1, 0, n)
			slice2 := other.Slice(0, o, j, j+1)
			value := slice1.DotProduct(slice2)
			result.SetValue(i, j, value)
		}
	}

	return result
}

//Calculates dot product of two vectors
func (mat *Matrix) DotProduct(other *Matrix) float64 {
	m, n := mat.Dimensions()
	o, p := other.Dimensions()
	vect1 := mat
	vect2 := other

	//Rotate matrices if needed
	if m == 1 {
		vect1 = mat.Rotate(true)
	}
	if p == 1 {
		vect2 = other.Rotate(false)
	}

	//Validate matrices are vectors of same length
	m, n = vect1.Dimensions()
	o, p = vect2.Dimensions()
	if n != 1 || o != 1 {
		m, n := mat.Dimensions()
		o, p := other.Dimensions()
		panic(fmt.Sprintf("Cannot compute dot product of non vector matrices. Matrices are of size (%v, %v), "+
			"and (%v, %v)", m, n, o, p))
	}
	if m != p {
		panic(fmt.Sprintf("Cannot compute dot product of vectors of different lengths. Vectors are of length "+
			"%v and %v", m, o))
	}

	//Inputs are good, we can calculate dot product
	result := 0.0

	for i := 0; i < m; i++ {
		result += vect1.GetValue(i, 0) * vect2.GetValue(0, i)
	}

	return result
}

//Rotates a matrix such that an M x N matrix becomes an N x M matrix
//The matrix is rotated 90 degrees
func (mat *Matrix) Rotate(clockwise bool) *Matrix {
	//Create correctly sized matrix
	m, n := mat.Dimensions()
	rotated := Matrix{}.NewMatrix(n, m)

	//Copy values
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if clockwise {
				rotated.SetValue(j, m-i-1, mat.GetValue(i, j))
			} else {
				rotated.SetValue(n-j-1, i, mat.GetValue(i, j))
			}
		}
	}

	return rotated
}

//Add two matrices
func (mat *Matrix) Add(other *Matrix) *Matrix {
	//Verify inputs
	m, n := mat.Dimensions()
	o, p := other.Dimensions()
	if m != o || n != p {
		panic(fmt.Sprintf("Cannot add matrices of different sizes. Matrixes are of sizes (%v, %v), and (%v, %v)",
			m, n, o, p))
	}

	result := Matrix{}.NewMatrix(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			result.SetValue(i, j, mat.GetValue(i, j)+other.GetValue(i, j))
		}
	}

	return result

}
