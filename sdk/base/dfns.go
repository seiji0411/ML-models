package base

func DSigmoid(x float64) float64 {
	return Sigmoid(x) * (1.0 - Sigmoid(x))
}
