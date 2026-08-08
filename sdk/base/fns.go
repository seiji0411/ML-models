package base

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// Dot ...
func Dot(w, x *mat.VecDense) float64 {
	return mat.Dot(w, x)
}

// Sigmoid ...
func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}
