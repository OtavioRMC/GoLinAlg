package vector

import (
	"math"
	"fmt"
)

type Vector struct {
	vector_data []float64
	Dimension   int
}

func NewVector() *Vector {
	return &Vector{}
}

func NewVectorWithDimensions(numDims int) *Vector {
	return &Vector{vector_data: make([]float64, numDims), Dimension: numDims}
}

func NewVectorWithInputData(inputData []float64) *Vector {
	return &Vector{vector_data: inputData, Dimension: len(inputData)}
}

func (v *Vector) GetNumDims() int {
	return v.Dimension
}

func (v *Vector) SetDim(numDims int) {
	if numDims < 0 {
		return
	}
	v.Dimension = numDims
	v.vector_data = make([]float64, numDims)
}

func (v *Vector) GetElement(index int) float64 {
	if index < 0 || index >= v.Dimension {
		return 0
	}
	return v.vector_data[index]
}

func (v *Vector) EuclideanNorm() float64 {
	sum := 0.0
	for _, val := range v.vector_data {
		sum += val * val
	}
	return math.Sqrt(sum)
}

func (v *Vector) NormalizedCopyVector() *Vector {
	norm := v.EuclideanNorm()
	
	if norm == 0 {
		return &Vector{}
	}

	normalizedData := make([]float64, v.Dimension)
	for i,val := range v.vector_data {
		normalizedData[i] = val / norm
	}

	return &Vector{vector_data: normalizedData, Dimension: v.Dimension}

}

func (v *Vector) NormalizeVector() {
	norm := v.EuclideanNorm()

	if norm != 0 {
		for i, val := range v.vector_data {
			v.vector_data[i] = val / norm
		}
	}
}

func (v *Vector) Add(vectorA *Vector) *Vector {
	if v.Dimension != vectorA.Dimension {
		return &Vector{}
	}
	resultVector := make([]float64, v.Dimension)
	for i := 0; i < v.Dimension; i++ {
		resultVector[i] = v.vector_data[i] + vectorA.vector_data[i]
	}
	return &Vector{vector_data: resultVector, Dimension: v.Dimension}
}

func (v *Vector) Subtract(vectorB *Vector) *Vector {
	if v.Dimension != vectorB.Dimension {
		return &Vector{}
	}
	resultVector := make([]float64, v.Dimension)
	for i := 0; i < v.Dimension; i++ {
		resultVector[i] = v.vector_data[i] - vectorB.vector_data[i]
	}
	return &Vector{vector_data: resultVector, Dimension: v.Dimension}
}

func DotProduct(vectorA, vectorB *Vector) float64 {
	if vectorA.Dimension != vectorB.Dimension {
		return 0.0
	}
	dotProduct := 0.0
	for i := 0; i < vectorA.Dimension; i++ {
		dotProduct += vectorA.vector_data[i] * vectorB.vector_data[i]
	}
	return dotProduct
}

func CrossProduct(vectorA, vectorB *Vector) *Vector {
	if vectorA.Dimension != 3 || vectorB.Dimension != 3 {
		return &Vector{}
	}
	resultVector := []float64{
		vectorA.vector_data[1]*vectorB.vector_data[2] - vectorA.vector_data[2]*vectorB.vector_data[1],
		vectorA.vector_data[2]*vectorB.vector_data[0] - vectorA.vector_data[0]*vectorB.vector_data[2],
		vectorA.vector_data[0]*vectorB.vector_data[1] - vectorA.vector_data[1]*vectorB.vector_data[0],
	}
	return &Vector{vector_data: resultVector, Dimension: 3}
}


func (v *Vector) Print() string {
	str := "["
	for i, val := range v.vector_data {
		str += fmt.Sprintf("%.2f", val)
		if i < v.Dimension-1 {
			str += ", "
		}
	}
	str += "]"
	return str
}

func (v *Vector) SetData(data []float64) {
	v.vector_data = data
	v.Dimension = len(data)
}
