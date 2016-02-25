package utils

import (
"github.com/gonum/matrix/mat64"
"math"
)

// This package contains all the mathematical computations needed by the regression algorithms


func GradientDescent(X, Y, theta *mat64.Dense, alpha float64, iter int) {
	m, _ := Y.Caps()
	H := mat64.NewDense(m, 1, nil)
	temp := mat64.NewDense(1, 2, nil)
	C := mat64.NewDense(m, 1, nil)
	for i := 0; i <= iter; i++ {
		H.Mul(X, theta.T())
		C.Apply(func(i, j int, v float64) float64 { return -v }, Y)
		H.Add(H, C)
		for j := 0; j < 2; j++ {
			temp.Set(0, j, theta.At(0, j)-((alpha/float64(m))*mat64.Sum(dotProduct(H, X.ColView(j)))))
		}
		for j := 0; j < 2; j++ {
			theta.Set(0, j, temp.At(0, j))
		}
		computeCost(X, Y, theta)
	}

}

func dotProduct(X *mat64.Dense, Y *mat64.Vector) *mat64.Dense {

	l, _ := X.Dims()
	P := mat64.NewDense(l, 1, nil)
	for i := 0; i < l; i++ {
		P.Set(i, 0, X.At(i, 0)*Y.At(i, 0))
	}
	return P
}

func computeCost(X, Y, theta *mat64.Dense) (cost float64) {
	m, _ := Y.Caps()
	H := mat64.NewDense(m, 1, nil)
	C := mat64.NewDense(m, 1, nil)
	H.Mul(X, theta.T())
	C.Apply(func(i, j int, v float64) float64 { return -v }, Y)
	H.Add(H, C)

	H.Apply(func(i, j int, v float64) float64 { return math.Pow(v, 2) }, H)
	cost = (1 / (2 * float64(m)) * mat64.Sum(H))
	//fmt.Println("computed cost ", cost)
	return
}