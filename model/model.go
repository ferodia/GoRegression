package model

import (
	"github.com/ferodia/GoRegression/linear"
	"github.com/ferodia/GoRegression/logistic"
)

const (
	Linear = iota
	Logistic
)

type Model interface {
	Initialise(x [][]float64, y []float64) error
	Train(iter int, alpha float64) error
	Predict(x [][]float64) []float64
}

func NewModel(n int) Model {
	switch n {
	case Linear:
		return &(linear.Linear{})
	case Logistic:
		return &(logistic.Logistic{})
	}
	return nil
}
