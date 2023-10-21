package matrix

import(
	"math"
	"errors"
	"fmt"
)

type Matrix struct {

	Rows , Cols int
	matrix_data []float64

}

func NewMatrix() *Matrix {
	return &Matrix{}
}

func NewMatrixWithSize(nRows, nCols int) *Matrix {

matrixData := make([]float64, nRows*nCols)
return &Matrix{
	Rows: nRows,
	Cols: nCols,
	matrix_data: matrixData,
}

}

func NewMatrixWithData(nRows, nCols int, InputData []float64) *Matrix {

	matrixData := make([]float64, nRows*nCols)

	copy(matrixData,InputData)
		
	return &Matrix{
			Rows: nRows,
			Cols: nCols,
			matrix_data: matrixData,
	
	}

}



func (m *Matrix) Resize(numRows, numCols int){
	m.Rows = numRows
	m.Cols = numCols
	m.matrix_data = make([]float64, numCols*numRows)
}

func (m *Matrix) SetToIdentity(){

	if m.Cols != m.Rows{
		return
	}

	for i := 0 ; i < m.Rows; i++ {
		for j:= 0; j < m.Cols; j++{
			if i == j {
				m.matrix_data[i * m.Cols + j] = 1.0 
			} else {
				m.matrix_data[i * m.Cols + j] = 0.0
			}
		}
	}
}

func (m *Matrix) GetElement(nRows, nCols int) float64{
	index := nRows * m.Cols + nCols
	if nRows < 0 || nRows >= m.Rows || nCols < 0 || nCols >= m.Cols{
		return 0.0
	}

	return m.matrix_data[index] 
}

func (m *Matrix) GetNumRows() int {
	return m.Rows
}

func (m *Matrix) GetNumCols() int {
	return m.Cols
}

func (m *Matrix) SetElement(nRows, nCols int, elementValue float64 ) error {
	index := nRows * m.Cols + nCols
	if nRows < 0 || nRows >= m.Rows || nCols < 0 || nCols >= m.Cols{
		return errors.New("Index out of Bounds")
	}
	m.matrix_data[index] = elementValue
	return nil
}

func (m *Matrix) Compare(matrix1 *Matrix, tolerance float64) bool {

if  matrix1.Rows != m.Rows || matrix1.Cols != m.Cols {
	return false
}

var cumulativeSum float64
numElements := m.Rows * m.Cols
for i:= 0; i < numElements ; i++ {
	element1 := float64(matrix1.matrix_data[i])
	element2 := float64(m.matrix_data[i])
	cumulativeSum += (element1 - element2) * (element1 - element2)
}

finalValue := math.Sqrt(cumulativeSum / float64(matrix1.Rows*matrix1.Cols - 1))
if finalValue < tolerance{
	return true
} 
return false

}

func (m *Matrix) IsSquare() bool {
	if m.Rows == m.Cols {
		return true
	}
return false
}

func (m *Matrix) FindSubMatrix(rowToRemove, colToRemove int) (*Matrix, error) {
	if rowToRemove >= m.Rows || colToRemove >= m.Cols {
			return nil, errors.New("Index out of bounds")
	}

	subMatrix := NewMatrixWithSize(m.Rows-1, m.Cols-1)
	subMatrixRow := 0
	for i := 0; i < m.Rows; i++ {
			if i == rowToRemove {
				continue
			}
			subMatrixCol := 0
			for j:= 0; j < m.Cols; j++{
				if j == colToRemove {
					continue
				}
				subMatrix.matrix_data[subMatrixRow * subMatrix.Cols + subMatrixCol] = m.GetElement(i,j)
				subMatrixCol++
			}
			subMatrixRow++
	}
	return subMatrix, nil
}


func (m *Matrix) Determinant() (float64, error) {
	if !m.IsSquare() {
			return 0, errors.New("Matrix is not square")
	}
	
	if m.Rows == 2 {
			return m.matrix_data[0]*m.matrix_data[3] - m.matrix_data[1]*m.matrix_data[2], nil
	} else {
			sum := 0.0
			for i := 0; i < m.Cols; i++ {
					subMatrix, _ := m.FindSubMatrix(0, i)
					det, _ := subMatrix.Determinant()
					sum += math.Pow(-1, float64(i)) * m.GetElement(0, i) * det
			}
			return sum, nil
	}
}


func (m *Matrix) MatrixPrint() {
	for i := 0; i < m.Rows; i++ {
		fmt.Print("[")
		for j := 0; j < m.Cols; j++ {
			if j != m.Cols-1 {
				fmt.Printf("%v , ", m.GetElement(i, j))
			} else {
				fmt.Printf("%v", m.GetElement(i, j)) // no comma for the last element
			}
		}
		fmt.Println("]")
	}
}

func (m *Matrix) MatrixSum(matrix1 *Matrix) (*Matrix, error) {

	if m.Cols == matrix1.Cols && m.Rows == matrix1.Rows{
		
		resultMatrix := NewMatrixWithSize(m.Rows, m.Cols)
		
		for i := 0;	i < m.Rows; i++{
			for j:= 0; j < m.Cols; j++{
				element := m.GetElement(i,j) + matrix1.GetElement(i,j)
				resultMatrix.SetElement(i,j,element)
			}
		}
		return resultMatrix, nil
	
	} else {
		return nil, errors.New("Impossible to Sum matrices of diferent dimensions")	
	}
}

func (m *Matrix) MatrixSubtract(matrix1 *Matrix) (*Matrix, error) {

	if m.Cols == matrix1.Cols && m.Rows == matrix1.Rows{
		
		resultMatrix := NewMatrixWithSize(m.Rows, m.Cols)
		
		for i := 0;	i < m.Rows; i++{
			for j:= 0; j < m.Cols; j++{
				element := m.GetElement(i,j) - matrix1.GetElement(i,j)
				resultMatrix.SetElement(i,j,element)
			}
		}
		return resultMatrix, nil
	
	} else {
		return nil, errors.New("Impossible to Sum matrices of diferent dimensions")	
	}
	
}

func (m *Matrix) HadamardProduct(matrix1 *Matrix) (*Matrix,error) {

if m.Cols == matrix1.Cols && m.Rows == matrix1.Rows {

	resultMatrix := NewMatrixWithSize(m.Cols, m.Rows)

	for i := 0; i < m.Cols; i++{
		for j:= 0; j < m.Rows; j++{
			product := m.GetElement(i,j) * matrix1.GetElement(i,j)
			resultMatrix.SetElement(i,j,product)
		}
	} 
		return resultMatrix, nil
}

return nil, errors.New("Matrix dimensions do not match for Hadamard product")


}

func (m *Matrix) MatrixMultiply(matrix1 *Matrix) (*Matrix, error) {
	if m.Cols != matrix1.Rows {
			return nil, errors.New("Matrix dimensions are not compatible for matrix multiplication")
	}

	resultMatrix := NewMatrixWithSize(m.Rows, matrix1.Cols)

	for i := 0; i < m.Rows; i++ {
			for j := 0; j < matrix1.Cols; j++ {
					var sum float64
					for k := 0; k < m.Cols; k++ {
							sum += m.GetElement(i, k) * matrix1.GetElement(k, j)
					}
					resultMatrix.SetElement(i, j, sum)
			}
	}

	return resultMatrix, nil
}
