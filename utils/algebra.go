package utils

import (
"github.com/gonum/matrix/mat64"
"math"
)

// This package contains all the mathematical computations needed by the regression algorithms


// Apply simoid function on every element of the matrix
func SigmoidMatrix(X *mat64.Dense) ( *mat64.Dense) {
	sigmat := &mat64.Dense{}
	sigmat.Apply(func(i, j int, v float64) float64 { return sigmoid(v) },X)
	return sigmat
}

// Subtract same value from all elements of the matrix
func Subtract(x float64, A * mat64.Dense) (*mat64.Dense)  {
	S := &mat64.Dense{}
	S.Apply(func(i, j int, v float64) float64 { return x-v }, A)
	return S
}

// Compute the logarithm of each element in the matrix
func Log(X *mat64.Dense) (*mat64.Dense) {
	S := &mat64.Dense{}
	S.Apply(func(i, j int, v float64) float64 { return math.Log(v) }, X)
	return S
}

func sigmoid(x float64) (z float64) {
	return(1/(1+math.Exp(-x)))
}


func Accuracy(y, pred []float64) float64 {
	acc := 0.0
	for index, value := range y {
		if value == pred[index]{
			acc++
		}
	}
	return acc/float64(len(y))
}
