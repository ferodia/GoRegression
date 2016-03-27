package logistic

import (
	"github.com/ferodia/GoRegression/utils"
	"github.com/gonum/matrix/mat64"
	"fmt"
)

type Logistic struct {
	X     *mat64.Dense
	Y     *mat64.Dense
	Theta *mat64.Dense
}

func (lg *Logistic) Initialise(x [][]float64, y []float64) error {

	m := len(y)
	features := len(x[0])
	lg.Theta = mat64.NewDense(1, 3, nil) // start with 0 values for theta
	lg.X = mat64.NewDense(m, 3, nil)


	for i := 0; i < m; i++ {
		lg.X.Set(i, 0, 1)

		for j := 1; j < features+1 ; j++ {
			lg.X.Set(i,j,x[i][j-1])

		}

	}

	lg.Y = mat64.NewDense(len(y), 1, y)

	return nil
}

func (lg *Logistic) Train(iterations int, alpha float64) error {

	//gradient descent
	gradientDescent(lg.X, lg.Y, lg.Theta, alpha, iterations)

	return nil
}

func (lg *Logistic) Predict(x [][]float64) (predictions []float64) {

	H := hypothesis(lg.X, lg.Theta)
	fmt.Println("prediction hypothesis ",H)
	m,_ := H.Dims()
	for i := 0; i < m; i++ {
		if H.At(i,0) >= 0.65 {
			predictions = append(predictions, 1)
		} else {
			predictions = append(predictions, 0)
		}

	}
	return
}

func gradientDescent(X, Y, theta *mat64.Dense, alpha float64, iter int) {
	m,_ := Y.Dims()
	H := &mat64.Dense{}
	temp := &mat64.Dense{}
	for i := 0; i <= iter; i++ {
		H.Sub(hypothesis(X,theta),Y)
		//fmt.Println("gradient decent hypothesis for iteration ",i, H)
		temp.Mul(H.T(),X)
		for j := 0; j < 3; j++ {
			theta.Set(0,j,theta.At(0,j)-((temp.At(0,j)*alpha)/float64(m)))
		}
		cost := computeCost(X, Y, theta)
		fmt.Println(" cost is ",cost)
	}

}

func computeCost(X, Y, theta *mat64.Dense) float64 {

	m, _ := Y.Caps()
	H := hypothesis(X, theta)
	fmt.Println("hypothesis ",H)
	A, B :=  &mat64.Dense{}, &mat64.Dense{}
	A.MulElem(utils.Subtract(0, Y), utils.Log(H))

	B.MulElem(utils.Subtract(1, Y), utils.Log(utils.Subtract(1, H)))
	A.Sub(A, B)

	return mat64.Sum(A)/ (float64(m))
}


func gradient(j int, X, Y, Theta *mat64.Dense) (float64) {
	H := hypothesis(X, Theta)
	//fmt.Println("hypothesis matrix",H)
	H.Sub(H, Y)
	H.MulElem(H, X.ColView(j))
	//fmt.Println(mat64.Sum(H))
	return mat64.Sum(H)

}

func hypothesis(X, Theta mat64.Matrix) ( *mat64.Dense) {
	H := &mat64.Dense{}
	//fmt.Println("theta ",Theta)
	H.Mul(X,Theta.T())
	//fmt.Println("inside hypothesis ",H)
	return utils.SigmoidMatrix(H)
}


