package linear

import (
	"github.com/ferodia/GoRegression/utils"
	"github.com/gonum/matrix/mat64"
	"math"
)

type Linear struct {
	x     *mat64.Dense
	y     *mat64.Dense
	theta *mat64.Dense
}

func (ln *Linear) Initialise(x [][]float64, y []float64) error {

	m := len(y) // number of examples in training set
	features := len(x[0]) // number of features
	ln.theta = mat64.NewDense(1,features+1, nil) // start with 0 values for theta
	ln.x = mat64.NewDense(m, features+1, nil)
	for i := 0; i < m; i++ {
		ln.x.Set(i, 0, 1.0)

		for j := 1; j < features +1; j++ {
			ln.x.Set(i, j, x[i][j-1])
		}

	}
	ln.y = mat64.NewDense(m, 1, y)

	return nil
}

func (ln *Linear) Train(iterations int, alpha float64) error {

	//gradient descent
	gradientDescent(ln.x, ln.y, ln.theta, alpha, iterations)

	return nil
}

func (ln *Linear) Predict(x [][]float64) (predictions []float64) {
	m := len(x)
	features := len(x[0])
	for i := 0; i < m; i++ {
		pred := ln.theta.At(0, 0)
		for j := 1; j < features+1;j++ {
			pred += x[i][j-1]*ln.theta.At(0, j)
		}
		predictions = append(predictions, pred)
	}
	return
}

func gradientDescent(x, y, theta *mat64.Dense, alpha float64, iter int) {
	m, features  := x.Caps()
	H := mat64.NewDense(m, 1, nil)
	temp := mat64.NewDense(1, features, nil)
	for i := 0; i <= iter; i++ {
		H.Mul(x, theta.T())
		H.Sub(H,y)
		for j := 0; j < features; j++ {
			temp.Set(0, j, theta.At(0, j)-((alpha/float64(m))*mat64.Sum(utils.DotProduct(H, x.ColView(j)))))
		}
		for j := 0; j < features; j++ {
			theta.Set(0, j, temp.At(0, j))
		}
		computeCost(x, y, theta)
	}

}

func computeCost(x, y, theta *mat64.Dense) float64 {
	m, _  := x.Caps()
	H := mat64.NewDense(m, 1, nil)
	H.Mul(x, theta.T())
	H.Sub(H,y)
	H.Apply(func(i, j int, v float64) float64 { return math.Pow(v, 2) }, H)
	return (1 / (2 * float64(m)) * mat64.Sum(H))
}
