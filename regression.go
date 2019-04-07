//Put regression code here
package regression

import (
	"errors"
	"fmt"
)

/*
Take in vectors: x, y
Optimize for SumofSquares
Compare w/ Py output
*/

//LinReg contains linear regression params
type LinReg struct {
	intercept float64
	coefs     []float64
}

//Compute errors of `preds` vs `truth`
func GetErrors(preds, truth []float64) ([]float64, error) {
	//check that preds and truth same length
	if len(preds) != len(truth) || len(preds) == 0 {
		return []float64{}, errors.New(fmt.Sprintf("GetErrors: can't subtract slices truth len %d, and preds len %d", len(truth), len(preds)))
	}
	//iterate over truth, and compute v preds
	outputDiffs := []float64{}
	for ind, t := range truth {
		fmt.Println(ind, t)
		diff := t - preds[ind]
		outputDiffs = append(outputDiffs, diff)
	}
	return outputDiffs, nil
}

//SquareVector squares the items in x
func SquareVector(x []float64) ([]float64, error) {
	if len(x) == 0 {
		errors.New(fmt.Sprintf("SquareVector: cannot work with input of len %d", len(x)))
	}
	outputSquares := make([]float64, 0)
	for _, number := range x {
		outputSquares = append(outputSquares, number*number)
	}
	return outputSquares, nil

}

func ComputeGradient() (float64, error) {
	return 1, nil
}

//SumVector sums all numbers in `x`
func SumVector(x []float64) (float64, error) {
	//add stuff up
	var sum float64
	for _, number := range x {
		sum += number
	}
	return sum, nil
}

func ComputeVectorMean(x []float64) (float64, error) {
	//sum
	vectorsum, _ := SumVector(x)
	mean := vectorsum / float64(len(x))
	return mean, nil
}

//# calculate mean
//m = sum(results) / len(results)

//# calculate variance using a list comprehension
//var_res = sum((xi - m) ** 2 for xi in results) / len(results)

//ComputeVectorVariance returns variance for x
func ComputeVectorVariance(x []float64) (float64, error) {
	diffs := make([]float64, len(x))
	meanx, _ := ComputeVectorMean(x)
	//subtract meanx from each item in x, this we can outsource
	for ind, number := range x {
		diff := number - meanx
		diffs[ind] = diff * diff
	}
	variance, _ := SumVector(diffs)
	return variance / float64(len(x)-1), nil
}

func AddtoVector(x []float64, add float64) ([]float64, error) {
	outputVector := make([]float64, len(x))
	for ind, number := range x {
		result := number + add
		fmt.Println("number", number, "out", result)
		outputVector[ind] = result
	}
	return outputVector, nil
}

//Compute the covariance between two vectors
/*
Cov(X,Y) = Σ
E(
	(X-μ)
	E(Y-ν)
	) / n-1 where:
X is a random variable
E(X) = μ is the expected value (the mean) of the random variable X and
E(Y) = ν is the expected value (the mean) of the random variable Y
n = the number of items in the data set


from numpy
sum items in m along colums
>>> m -= np.sum(m * w, axis=1, keepdims=True) / v1
multiply the summed items with transpose
>>> cov = np.dot(m * w, m.T) * v1 / (v1**2 - ddof * v2)

Calculate covariance for the following data set:
x: 2.1, 2.5, 3.6, 4.0 (mean = 3.1)
y: 8, 10, 12, 14 (mean = 11)

Substitute the values into the formula and solve:
Cov(X,Y) = ΣE((X-μ)(Y-ν)) / n-1
= (2.1-3.1)(8-11)+(2.5-3.1)(10-11)+(3.6-3.1)(12-11)+(4.0-3.1)(14-11) /(4-1)
= (-1)(-3) + (-0.6)(-1)+(.5)(1)+(0.9)(3) / 3
= 3 + 0.6 + .5 + 2.7 / 3
= 6.8/3
= 2.267
*/
// ComputeVectorCovariance returns covariance between x and y (uses n-1)
func ComputeVectorCovariance(x, y []float64) (float64, error) {
	if len(x) != len(y) {
		return 0, fmt.Errorf(
			"ComputeVectorCovariance: input vectors are len %d , %d, needs to be ==", len(x), len(y))
	}
	meanx, _ := ComputeVectorMean(x)
	meany, _ := ComputeVectorMean(y)
	outvector := make([]float64, len(x))
	//make vector of x-meanz, y-meany
	for ind, number := range x {
		resx := number - meanx
		resy := y[ind] - meany
		outvector[ind] = resx * resy
	}
	vecTotal, _ := SumVector(outvector)
	covar := vecTotal / float64(len(outvector)-1)
	return covar, nil
}

//Regressor fits LinReg struct on input values x and output y
func Regressor(x [][]float64, y []float64) (LinReg, error) {
	//takes in x, y returns fitted LinReg
	meany, _ := ComputeVectorMean(y)
	weights := make([]float64, len(x))
	var intercept float64

	for ind, vec := range x {
		variance, err := ComputeVectorVariance(vec)
		covar, err := ComputeVectorCovariance(vec, y)
		if err != nil {
			fmt.Errorf("Regressor: Covariance returns %s", err)
			return LinReg{}, err
		}
		weight := (covar / variance)
		weights[ind] = weight
		meanx, _ := ComputeVectorMean(vec)
		intercept = meany - (weight * meanx)
		fmt.Printf("weight is %e", weight)
		fmt.Printf("intercept is %e", intercept)
	}
	return LinReg{intercept, weights}, nil
}

// Coefs
// B1 = sum((x(i) - mean(x)) * (y(i) - mean(y))) / sum( (x(i) - mean(x))^2 )
// B0 = mean(y) - B1 * mean(x)
// B1 = covariance(x, y) / variance(x)

//check weights for fit
