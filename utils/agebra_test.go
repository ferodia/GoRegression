package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSigmoidMatrix(t *testing.T) {
	for _, test := range sigmoidCases {
		m,n := test.X.Dims()
		sigmoidMatrix := SigmoidMatrix(test.X)
		for i:=0;i<m;i++ {
			for j:=0;j<n;j++ {
				assert.Equal(t,test.sigmoid,sigmoidMatrix.At(i,j))
			}
		}
	}
}

func TestSubtract(t *testing.T){
	for _, test := range sigmoidCases {
		m,n := test.X.Dims()
		sub := Subtract(test.elem, test.X)
		for i:=0;i<m;i++ {
			for j:=0;j<n;j++ {
				assert.Equal(t,0.0,sub.At(i,j))
			}
		}
	}
}

func TestLog(t *testing.T){
	for _, test := range logCases {
		m,n := test.X.Dims()
		log := Log(test.X)
		for i:=0;i<m;i++ {
			for j:=0;j<n;j++ {
				assert.Equal(t,test.log,log.At(i,j))
			}
		}
	}
}

func TestSigmoid(t *testing.T) {
	for _, test := range sigmoidCases {
		assert.Equal(t,test.sigmoid,sigmoid(test.elem))
	}
}
