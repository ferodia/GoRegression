package regression

import "github.com/gonum/matrix/mat64"

var cases = []struct {
	X        *mat64.Dense
	Y        *mat64.Dense
	theta    *mat64.Dense
	pred     []float64
	cost     float64
	product  *mat64.Dense
	minTheta *mat64.Dense
}{
	{
		mat64.NewDense(2, 2, []float64{1, 2, 3, 4}),
		mat64.NewDense(2, 1, []float64{-2, 2}),
		mat64.NewDense(1, 2, []float64{3.5, -1}),
		[]float64{1.5, -0.5},
		8.125,
		mat64.NewDense(2, 1, []float64{-2, 6}),
		mat64.NewDense(1, 2, []float64{1.176047178559446e+18, 1.66881672138228e+18})},
	{
		mat64.NewDense(3, 2, []float64{7, 1, 0, 9, 7, 9}),
		mat64.NewDense(3, 1, []float64{2, -1, 10}),
		mat64.NewDense(1, 2, []float64{5, 4}),
		[]float64{9, 41, 41},
		1076.5,
		mat64.NewDense(3, 1, []float64{14, 0, 70}),
		mat64.NewDense(1, 2, []float64{6.529437753443026e+29, 1.0230392581676892e+30})},
	{
		mat64.NewDense(5, 2, []float64{-8, 0, 3.7, 4, 7.9, 8, -9, 9, 6, 8}),
		mat64.NewDense(5, 1, []float64{0, 7, 8, 9, 10}),
		mat64.NewDense(1, 2, []float64{2, -1.5}),
		[]float64{2, -4, -10, -11.5, -10},
		204.525,
		mat64.NewDense(5, 1, []float64{0, 25.900000000000002, 63.2, -81, 60}),
		mat64.NewDense(1, 2, []float64{7.221929579900366e+26, 4.4431282352081704e+26})},
}

var RunCases = []struct {
	X           []float64
	Y           []float64
	predictions []float64
}{
	{
		[]float64{3, 1},
		[]float64{2, 4},
		[]float64{2.0851778560547225, 3.794362464698826}},
	{
		[]float64{4, 1, 2, 0},
		[]float64{0, 3, 0, 3.5},
		[]float64{-0.5233604175890236, 2.3351766656956015, 1.382330971267393, 3.2880223601238097}},

	{
		[]float64{-8, 0, 3.7, 4, 3, 2, 1},
		[]float64{2.3, 1, 3, 6, 6, 3, 9},
		[]float64{2.1878650547792198, 4.13080562803636, 5.029415643167788, 5.10227591466493, 4.859408343007788, 4.616540771350645, 4.373673199693503}},
}
