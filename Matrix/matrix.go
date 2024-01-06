package matrix

import(
	"math"
	"errors"
	"fmt"
)

type Matrix struct {

	rows , cols int
	matrixData []float64

}

func NewMatrix() *Matrix {
	return &Matrix{}
}

func NewMatrixWithSize(nRows, nCols int) *Matrix {

matrix_data := make([]float64, nRows*nCols)
return &Matrix{
	rows: nRows,
	cols: nCols,
	matrixData: matrix_data,
}

}

func NewMatrixWithData(nRows, nCols int, InputData []float64) *Matrix {

	matrixData := make([]float64, nRows*nCols)

	copy(matrixData,InputData)
		
	return &Matrix{
			rows: nRows,
			cols: nCols,
			matrixData: matrixData,
	
	}

}


func (m *Matrix) Resize(numRows, numCols int){
	m.rows = numRows
	m.cols = numCols
	m.matrixData = make([]float64, numCols*numRows)
}

func (m *Matrix) SetToIdentity(){

	if m.cols != m.rows{
		return
	}

	for i := 0 ; i < m.rows; i++ {
		for j:= 0; j < m.cols; j++{
			if i == j {
				m.matrixData[i * m.cols + j] = 1.0 
			} else {
				m.matrixData[i * m.cols + j] = 0.0
			}
		}
	}
}

func (m *Matrix) GetElement(nRows, nCols int) float64{
	index := nRows * m.cols + nCols
	if nRows < 0 || nRows >= m.rows || nCols < 0 || nCols >= m.cols{
		return 0.0
	}

	return m.matrixData[index] 
}

func (m *Matrix) GetNumRows() int {
	return m.rows
}

func (m *Matrix) GetNumCols() int {
	return m.cols
}

func (m *Matrix) SetElement(nRows, nCols int, elementValue float64 ) error {
	index := nRows * m.cols + nCols
	if nRows < 0 || nRows >= m.rows || nCols < 0 || nCols >= m.cols{
		return errors.New("Index out of Bounds")
	}
	m.matrixData[index] = elementValue
	return nil
}

func (m *Matrix) Compare(matrix1 *Matrix, tolerance float64) bool {

if  matrix1.rows != m.rows || matrix1.cols != m.cols {
	return false
}

var cumulativeSum float64
numElements := m.rows * m.cols
for i:= 0; i < numElements ; i++ {
	element1 := float64(matrix1.matrixData[i])
	element2 := float64(m.matrixData[i])
	cumulativeSum += (element1 - element2) * (element1 - element2)
}

finalValue := math.Sqrt(cumulativeSum / float64(matrix1.rows*matrix1.cols - 1))
if finalValue < tolerance{
	return true
} 
return false

}

func (m *Matrix) IsSquare() bool {
	if m.rows == m.cols {
		return true
	}
return false
}

func (m *Matrix) FindSubMatrix(rowToRemove, colToRemove int) (*Matrix, error) {
	if rowToRemove >= m.rows || colToRemove >= m.cols {
			return nil, errors.New("Index out of bounds")
	}

	subMatrix := NewMatrixWithSize(m.rows-1, m.cols-1)
	subMatrixRow := 0
	for i := 0; i < m.rows; i++ {
			if i == rowToRemove {
				continue
			}
			subMatrixCol := 0
			for j:= 0; j < m.cols; j++{
				if j == colToRemove {
					continue
				}
				subMatrix.matrixData[subMatrixRow * subMatrix.cols + subMatrixCol] = m.GetElement(i,j)
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
	
	if m.rows == 2 {
			return m.matrixData[0]*m.matrixData[3] - m.matrixData[1]*m.matrixData[2], nil
	} else {
			sum := 0.0
			for i := 0; i < m.cols; i++ {
					subMatrix, _ := m.FindSubMatrix(0, i)
					det, _ := subMatrix.Determinant()
					sum += math.Pow(-1, float64(i)) * m.GetElement(0, i) * det
			}
			return sum, nil
	}
}


func (m *Matrix) MatrixPrint() {
	for i := 0; i < m.rows; i++ {
		fmt.Print("[")
		for j := 0; j < m.cols; j++ {
			if j != m.cols-1 {
				fmt.Printf("%v , ", m.GetElement(i, j))
			} else {
				fmt.Printf("%v", m.GetElement(i, j)) // no comma for the last element
			}
		}
		fmt.Println("]")
	}
}

func (m *Matrix) MatrixSum(matrix1 *Matrix) (*Matrix, error) {

	if m.cols == matrix1.cols && m.rows == matrix1.rows{
		
		resultMatrix := NewMatrixWithSize(m.rows, m.cols)
		
		for i := 0;	i < m.rows; i++{
			for j:= 0; j < m.cols; j++{
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

	if m.cols == matrix1.cols && m.rows == matrix1.rows{
		
		resultMatrix := NewMatrixWithSize(m.rows, m.cols)
		
		for i := 0;	i < m.rows; i++{
			for j:= 0; j < m.cols; j++{
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

if m.cols == matrix1.cols && m.rows == matrix1.rows {

	resultMatrix := NewMatrixWithSize(m.cols, m.rows)

	for i := 0; i < m.cols; i++{
		for j:= 0; j < m.rows; j++{
			product := m.GetElement(i,j) * matrix1.GetElement(i,j)
			resultMatrix.SetElement(i,j,product)
		}
	} 
		return resultMatrix, nil
}

return nil, errors.New("Matrix dimensions do not match for Hadamard product")


}

func (m *Matrix) MatrixMultiply(matrix1 *Matrix) (*Matrix, error) {
	if m.cols != matrix1.rows {
			return nil, errors.New("Matrix dimensions are not compatible for matrix multiplication")
	}

	resultMatrix := NewMatrixWithSize(m.rows, matrix1.cols)

	for i := 0; i < m.rows; i++ {
			for j := 0; j < matrix1.cols; j++ {
					var sum float64
					for k := 0; k < m.cols; k++ {
							sum += m.GetElement(i, k) * matrix1.GetElement(k, j)
					}
					resultMatrix.SetElement(i, j, sum)
			}
	}

	return resultMatrix, nil
}
