package utils



type Model interface {

	Initialise(x, y []float64) error
	Train(iter int, alpha float64) error
	Predict(x []float64) []float64
}

