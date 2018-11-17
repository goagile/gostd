package functions

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 6, add(1, 2, 3))
	assert.Equal(t, 6.0, sum([]float64{ 1, 2, 3 }))
}

func add(args ...int) int {
	result := 0
	for _, arg := range args {
		result += arg
	}
	return result
}

func sum(x []float64) float64 {
	var result float64 = 0
	for i:=0; i<len(x); i++ {
		result += x[i]
	}
	return result
}