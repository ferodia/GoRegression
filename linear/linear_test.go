package linear

/*
import (
	"github.com/gonum/matrix/mat64"
	"reflect"
	"testing"
)

func TestProduct(t *testing.T) {
	for _, test := range cases {
		observed := product(test.X, test.Y.ColView(0))
		if !mat64.Equal(test.product, observed) {
			t.Fatalf("expected :\n%v\ngot:\n%v\n\n ", mat64.Formatted(test.product, mat64.Prefix(""), mat64.Excerpt(0)),
				mat64.Formatted(observed, mat64.Prefix(""), mat64.Excerpt(0)))
		}
	}
}

func TestComputeCost(t *testing.T) {
	for _, test := range cases {
		observed := computeCost(test.X, test.Y, test.theta)
		if test.cost != observed {
			t.Fatal("expected ", test.cost, "got ", observed)
		}
	}
}

func TestPredict(t *testing.T) {
	for _, test := range cases {
		observed := predict(test.X, test.theta)
		if !reflect.DeepEqual(test.pred, observed) {
			t.Fatal("expected ", test.pred, "got ", observed)
		}
	}
}

func TestGradientDecent(t *testing.T) {
	for _, test := range cases {
		gradientDescent(test.X, test.Y, test.theta, 1, 15)
		observed := test.theta
		if !mat64.Equal(test.minTheta, observed) {
			t.Fatalf("expected :\n%v\ngot:\n%v\n\n ", mat64.Formatted(test.minTheta, mat64.Prefix(""), mat64.Excerpt(0)),
				mat64.Formatted(observed, mat64.Prefix(""), mat64.Excerpt(0)))
		}
	}
}

func TestTrain(t *testing.T) {
	for _, test := range RunCases {
		observed := Train(test.X, test.Y)
		if !reflect.DeepEqual(test.predictions, observed) {
			t.Fatal("expected ", test.predictions, "got ", observed)
		}
	}
}

func BenchmarkTrain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range RunCases {
			Train(test.X, test.Y)
		}
	}
}
*/
