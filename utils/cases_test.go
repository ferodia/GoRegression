package utils

import "github.com/gonum/matrix/mat64"

var sigmoidCases = []struct {
	X        *mat64.Dense
	elem	 float64
	sigmoid  float64
}{
	{
		mat64.NewDense(2,2,[]float64{0,0,0,0}),
		0,
		0.5,

	},
	{
		mat64.NewDense(3,2,[]float64{1000000000,1000000000,1000000000,1000000000,1000000000,1000000000}),
		1000000000,
		1.0,

	},
	{
		mat64.NewDense(2,4,[]float64{-1000000000,-1000000000,-1000000000,-1000000000,-1000000000,-1000000000,
			-1000000000,-1000000000}),
		-1000000000,
		0.0,

	},

}


var logCases = []struct {
	X        *mat64.Dense
	log	 float64
}{
	{
		mat64.NewDense(2,2,[]float64{2.7183,2.7183,2.7183,2.7183}),
		1.0000066849139877,
	},
	{
		mat64.NewDense(2,2,[]float64{1,1,1,1}),
		0,

	},
	{
		mat64.NewDense(2,2,[]float64{1000000,1000000,1000000,1000000}),
		13.815510557964274,

	},
}
