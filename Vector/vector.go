package vector

import (
	"math"
	"fmt"
)


type Vector struct {
	vectorData []float64
	dimension   int
}


func NewVector() *Vector {
	return &Vector{}
}


func NewVectorWithDimensions(numDims int) *Vector {
	return &Vector{vectorData: make([]float64, numDims), dimension: numDims}
}


func NewVectorWithInputData(inputData []float64) *Vector {
	return &Vector{vectorData: inputData, dimension: len(inputData)}
}


func (v *Vector) GetNumDims() int {
	return v.dimension
}


func (v *Vector) SetDim(numDims int) {
	
	if numDims < 0 {
		return
	}
	
	v.dimension = numDims
	v.vectorData = make([]float64, numDims)
}


func (v *Vector) GetElement(index int) float64 {
	
	if index < 0 || index >= v.dimension {
		return 0
	}
	
	return v.vectorData[index]

}

func (v *Vector) EuclideanNorm() float64 {
	
	sum := 0.0
	
	for _, val := range v.vectorData {
		sum += val * val
	}
	
	return math.Sqrt(sum)

}


func (v *Vector) NormalizedCopyVector() *Vector {
	
	norm := v.EuclideanNorm()
	
	if norm == 0 {
		return &Vector{}
	}

	normalizedData := make([]float64, v.dimension)
	
	for i,val := range v.vectorData {
		normalizedData[i] = val / norm
	}

	return &Vector{vectorData: normalizedData, dimension: v.dimension}

}

func (v *Vector) NormalizeVector() {
	
	norm := v.EuclideanNorm()

	if norm != 0 {
		
		for i, val := range v.vectorData {
			v.vectorData[i] = val / norm
		}
	
	}
}

func (v *Vector) Add(vectorA *Vector) *Vector {
	
	if v.dimension != vectorA.dimension {
		return &Vector{}
	}
	
	resultVector := make([]float64, v.dimension)
	
	for i := 0; i < v.dimension; i++ {
		resultVector[i] = v.vectorData[i] + vectorA.vectorData[i]
	}
	
	return &Vector{vectorData: resultVector, dimension: v.dimension}

}

func (v *Vector) Subtract(vectorB *Vector) *Vector {
	
	if v.dimension != vectorB.dimension {
		return &Vector{}
	}
	
	resultVector := make([]float64, v.dimension)
	
	for i := 0; i < v.dimension; i++ {
	
		resultVector[i] = v.vectorData[i] - vectorB.vectorData[i]
	
	}
	
	return &Vector{vectorData: resultVector, dimension: v.dimension}
}

func DotProduct(vectorA, vectorB *Vector) float64 {
	
	if vectorA.dimension != vectorB.dimension {
		return 0.0
	}
	
	dotProduct := 0.0
	
	for i := 0; i < vectorA.dimension; i++ {
	
		dotProduct += vectorA.vectorData[i] * vectorB.vectorData[i]
	
	}
	return dotProduct
}

func CrossProduct(vectorA, vectorB *Vector) *Vector {
	
	if vectorA.dimension != 3 || vectorB.dimension != 3 {
		return &Vector{}
	
	}
	
	resultVector := []float64{
		
		vectorA.vectorData[1]*vectorB.vectorData[2] - vectorA.vectorData[2]*vectorB.vectorData[1],
		vectorA.vectorData[2]*vectorB.vectorData[0] - vectorA.vectorData[0]*vectorB.vectorData[2],
		vectorA.vectorData[0]*vectorB.vectorData[1] - vectorA.vectorData[1]*vectorB.vectorData[0],
	
	}
	
	return &Vector{vectorData: resultVector, dimension: 3}

}


func (v *Vector) Print() string {
	
	str := "["
	
	for i, val := range v.vectorData {
	
		str += fmt.Sprintf("%.2f", val)
		if i < v.dimension-1 {
			str += ", "
		}
	
	}
	str += "]"
	return str
}

func (v *Vector) SetData(data []float64) {

	v.vectorData = data
	v.dimension = len(data)

}
