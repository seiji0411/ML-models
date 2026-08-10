package regression

import (
	"fmt"

	"github.com/seiji0411/ml-models/sdk/base"
	"github.com/seiji0411/ml-models/sdk/data"
	"gonum.org/v1/gonum/mat"
)

// Linear provides API to build linear regression model.
// Linear regression is a statistical technique used to find the relationship between variables.
// In an ML context, linear regression finds the relationship between features and a label.
// https://developers.google.com/machine-learning/crash-course/linear-regression
//
// The model is defined by the following equation:
// y = mx + b
// where:
// y is the label
// m is the slope
// x is the feature
// b is the intercept
type Linear struct {
	Regresser
}

// Fit trains the model if there are enough data points
// present to run the regression. If training has already completed on provided
// data points then function returns an error.
func (r *Linear) Fit(pts []data.Point) error {
	fmt.Println("Linear#Fit ...")
	defer fmt.Println("Linear#Fitted")
	return r.fit(pts, func(w, x *mat.VecDense) float64 {
		return base.Dot(w, x)
	})
}
