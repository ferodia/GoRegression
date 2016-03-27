package utils

import (
"github.com/gonum/matrix/mat64"
"math"
)

// This package contains all the mathematical computations needed by the regression algorithms




func DotProduct(X *mat64.Dense, Y *mat64.Vector) *mat64.Dense {

	l, _ := X.Dims()
	P := mat64.NewDense(l, 1, nil)
	for i := 0; i < l; i++ {
		P.Set(i, 0, X.At(i, 0)*Y.At(i, 0))
	}
	return P
}

// Apply simoid function on every element of the matrix
func SigmoidMatrix(X *mat64.Dense) ( *mat64.Dense) {
	sigmat := &mat64.Dense{}
	sigmat.Apply(func(i, j int, v float64) float64 { return sigmoid(v) },X)
	return sigmat
}

func Subtract(x float64, A * mat64.Dense) (*mat64.Dense)  {
	S := &mat64.Dense{}
	S.Apply(func(i, j int, v float64) float64 { return x-v }, A)
	return S
}

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
	return acc/float64(len(y)
}
