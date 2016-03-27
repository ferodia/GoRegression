package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSigmoid(t *testing.T) {
	assert.Equal(t,0.5,sigmoid(0))
	assert.Equal(t,0.0,sigmoid(1000000000))
	assert.Equal(t,1.0,sigmoid(-1000000000000))
}
