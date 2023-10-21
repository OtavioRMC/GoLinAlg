# GoLinAlg
A linear algebra package for Go

# Index
* [Getting Started](#getting-started)
* [Vector](#vector)
   * [Vector Usage Examples](#usage)
   * [Creating a Vector](#creating-a-vector)
   * [Acessing vector information](#acessing-vector-information)
   * [Modifying a vector](#modifying-a-vector)
   * [Vector Normalization](#vector-normalization)
   * [Vector Operations](#vector-operations)
   * [Vector Usage Examples](#vector-usage)
* [Matrix](#matrix)
  * [Creating a Matrix](#creating-a-matrix)
  * [Modifying a Matrix](#modifying-a-matrix)
  * [Acessing Matrix Information](#acessing-matrix-information)
  * [Printing a Matrix](#printing-a-matrix)
  * [Matrix Operations](#matrix-operations)
  * [Matrix Usage Examples](#matrix-usage)

# Getting Started

To add Go Linear Algebra into your projects you can:

1. Clone the repository or download the source code.
```shell
  git clone https://github/OtavioRMC/GoLinAlg
```
1.1 Include the package in your Go file. For example:
   ```go
   package main

   import (
    "GoLinAlg/Matrix"
    "GoLinAlg/Vector"
   )
   ```

2. Add direclty:
```go
package main

import (
  "github.com/yourusername/GoLinAlg/Vector"
  "github.com/yourusername/GoLinAlg/Matrix"
)
``` 

# Vector

This Package allows you to create and manipulate vectors of floating-point numbers.

## Creating a Vector

The package provides several methods to create and work with vectors.

- `NewVector() *Vector`: Create an Empty Vector

- `NewVectorWithDimensions(num Dims int) *Vector`: Create a Vector with a Specific Dimension

- `NewVectorWithInputData(inputData []float64) *Vector`: Create a Vector with Input Data

## Acessing Vector Information

- `GetNumDims() int`: Get The Number of Dimensions

- `GetElement(index int) float64`: Get the Value of a Vector Element in a given index

## Modifying a Vector

- `SetDim(numDims int)`: Set the Dimension of a Vector

- `SetData(data []float64)`: This method set the data of a vector by providing a slice of float64 values. 

## Vector Normalization

- `EuclideanNorm() float64`: Calculate the Euclidean Norm of a Vector

- `NormalizedCopyVector() *Vector`: Create a Normalized Copy of a Vector

- `NormalizeVector()`: Normalize a Vector in Place

## Vector Operations

- `Add(vectorA *Vector) *Vector`: Add Vectors

- `Subtract(vectorA *Vector) *Vector`: Subtract Vectors

- `DotProduct(vectorA,VectorB *Vector) float64`: Calculate the Dot Product of Two Vectors

- `CrossProduct(vectorA, vectorB *Vector) *Vector`: Calculate the Cross Product of two 3D vectors.

- `Print() string`: Print the vector returning a string representation.

## Vector Usage
```Go
  // Initializing a new vector with 3 dimensions
  v1 := vector.NewVectorWithDimensions(3)
  fmt.Println(v1.Print()) // prints: [0.00, 0.00, 0.00]
		
  // Set dimensions and values
  v1.SetDim(5)
  fmt.Println(v1.Print()) // prints: [0.00, 0.00, 0.00, 0.00, 0.00]
		
  // Initializing a new vector with input data
  v2 := vector.NewVectorWithInputData([]float64{1, 2, 3, 4, 5})
  fmt.Println(v2.Print()) // prints: [1.00, 2.00, 3.00, 4.00, 5.00]

  // Get number of dimensions
  fmt.Println(v2.GetNumDims()) // prints: 5
		
  // Get element at index
  fmt.Println(v2.GetElement(2)) // prints: 3.00
		
  // Compute Euclidean norm
  fmt.Println(v2.EuclideanNorm()) // prints: 7.42
			
  // Create a normalized copy
  v3 := v2.NormalizedCopyVector()
  fmt.Println(v3.Print()) // prints: [0.13, 0.27, 0.40, 0.53, 0.67]

  // Normalize the original vector
  v2.NormalizeVector()
  fmt.Println(v2.Print()) // prints: [0.13, 0.27, 0.40, 0.53, 0.67]
			
  // Add two vectors
  v4 := vector.NewVectorWithInputData([]float64{1, 1, 1, 1, 1})
  v5 := v2.Add(v4)
  fmt.Println(v5.Print()) // prints: [1.13, 1.27, 1.40, 1.53, 1.67]
		
  // Subtract two vectors
  v6 := v5.Subtract(v4)
  fmt.Println(v6.Print()) // prints: [0.13, 0.27, 0.40, 0.53, 0.67]
		
  // Compute dot product
  dot := vector.DotProduct(v4, v5)
  fmt.Println(dot) // prints: 6.00

  // Compute cross product (only applicable for 3D vectors)
  v7 := vector.NewVectorWithInputData([]float64{1, 2, 3})
  v8 := vector.NewVectorWithInputData([]float64{4, 5, 6})
  v9 := vector.CrossProduct(v7, v8)
  fmt.Println(v9.Print()) // prints: [-3.00, 6.00, -3.00]
```

# Matrix

This Package allows you to create and manipulate matrices of floating-point numbers.

## Creating a Matrix

You can create a new matrix using one of the following constructors:

- `NewMatrix() *Matrix` : Create an empty matrix.

- `NewMatrixWithSize(nRows, nCols int) *Matrix`: Create a matrix with a specified number of rows and columns

- `NewMatrixWithData(nRows,nCols int, InputData []float64) * Matrix`: Create a matrix with specified data. 

## Modifying a Matrix

- `Resize(nRows, nCols int)`: Resize the matrix to have the specified number of rows and columns.

- `SetToIdentity()`: Set the matrix to the identity matrix. (Only works for square matrices)

- `SetElement(nRows, nCols int, elementValue float64) error`: Set the value of a specific element in the matrix.

## Acessing Matrix Information

- `GetElement(nRows,nCols int) float64`: Retrieve the value of a specific element in the matrix.

- `GetNumRows() int`: Get the number of rows in the matrix.

- `GetNumCols() int`: Get the number of colummns in the matrix.

- `IsSquare() bool`: Check if the matrix is square (Having the same number of rows and columns).

## Printing a Matrix

- `Print()`: Display the matrix's elements in a human-readable format.

## Matrix Operations

- `Compare(matrix1 * Matrix, tolerance float64) bool`: Compare two matrices for equality with a specified tolerance.

- `FindSubMatrix(rowToRemove, colToRemove int) (*Matrix,error)`: Find a submatrix by removing a specified row and column.

- `Determinant() (float64,error)`: Calculate the determinant of the matrix.

- `MatrixSum(matrix1 *Matrix) (*Matrix,error)`: Sum Two Matrices

- `MatrixSubtract(matrix1 *Matrix) (*Matrix,error)`: Subtract Two Matrices

- `HadamardProduct(matrix1 *Matrix) (*Matrix,error)`: Calculate the Hadard Product

- `MatrixMultiply(matrix1 *Matrix) (*Matrix,error)`: Multiply two matrices

##  Matrix Usage
```go
  // Create Matrix with determined size
  m := matrix.NewMatrixWithSize(2, 2)
	
  // Seting Element
  m.SetElement(0, 0, 1.0)
  m.MatrixPrint()

  // Getting Element
  fmt.Println(m.GetElement(0, 0))


  // Performing Matrix Addition
  m1 := matrix.NewMatrixWithSize(2, 2)
  m2 := matrix.NewMatrixWithSize(2, 2)
  m1.SetElement(0, 0, 1.0)
  m2.SetElement(0, 0, 2.0)
  m3, _ := m1.MatrixSum(m2)
  m3.MatrixPrint()


  // Matrix Multiplication
  m3, _ := m1.MatrixMultiply(m2)

  // Calculating Determinant
  m4 := matrix.NewMatrixWithSize(2, 2)
  m4.SetElement(0, 0, 1.0)
  m4.SetElement(0, 1, 2.0)
  m4.SetElement(1, 0, 3.0)
  m4.SetElement(1, 1, 4.0)
  det, _ := m4.Determinant()
  fmt.Prinln(det)

  m5 := matrix.NewMatrixWithSize(2, 2)
  m6 := matrix.NewMatrixWithSize(2, 2)
  m5.SetElement(0, 0, 1.0)
  m5.SetElement(0, 1, 2.0)
  m5.SetElement(1, 0, 3.0)
  m5.SetElement(1, 1, 4.0)
  m6.SetElement(0, 0, 5.0)
  m6.SetElement(0, 1, 6.0)
  m6.SetElement(1, 0, 7.0)
  m6.SetElement(1, 1, 8.0)
  result, _ := m5.HadamardProduct(m6)
  fmt.Println(result)

``` 
## Contributing

I welcome contributions to this project. If you're interested in contributing, here are some general guidelines to follow:

- Feel free to fork this repository.
- Make your changes or improvements.
- Submit a Pull Request with a clear description of your changes.
- I will review your Pull Request and provide feedback.

Thank you for considering contributing to GoLinAlg!
