package regression

import (
	"github.com/gonum/matrix/mat64"
	"github.com/ferodia/GoRegression/utils"
)



type Linear struct {
	X *mat64.Dense
	Y *mat64.Dense
	Theta *mat64.Dense
}

func (ln * Linear) Initialise(x,y []float64) error {

	ln.Theta = mat64.NewDense(1, 2, nil) // start with 0 values for theta
	ln.X = mat64.NewDense(len(x), 2, nil)
	for i := 0; i < len(y); i++ {
		ln.X.Set(i, 0, 1.0)
	}
	ln.X.SetCol(1, x)

	ln.Y = mat64.NewDense(len(y), 1, y)

	return nil
}

func (ln * Linear) Train(iterations int, alpha float64) error {

	//gradient descent
	utils.GradientDescent(ln.X, ln.Y, ln.Theta, alpha, iterations)

	return nil
}

func (ln * Linear) Predict(x []float64) (predictions []float64) {
	l := len(x)
	for i := 0; i < l; i++ {
		predictions = append(predictions, ln.Theta.At(0, 0)+(x[i]*ln.Theta.At(0, 1)))
	}
	return
}
