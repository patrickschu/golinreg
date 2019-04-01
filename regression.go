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

//ComputeVectorVariance returns vector of variances for x
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
Cov(X,Y) = Σ E((X-μ)E(Y-ν)) / n-1 where:
X is a random variable
E(X) = μ is the expected value (the mean) of the random variable X and
E(Y) = ν is the expected value (the mean) of the random variable Y
n = the number of items in the data set
*/

func ComputeVectorCovariance(x, y []float64) (float64, error) {
	meanx, _ := ComputeVectorMean(x)
	meany, _ := ComputeVectorMean(y)
	//make vector of x-meanz, y-meany
	diffx, _ := AddtoVector(x, -meanx)
	diffy, _ := AddtoVector(y, -meany)
	fmt.Println("diffx", diffx)
	fmt.Println("diffy", diffy)
	//get the mean
	meanDiffy, _ := ComputeVectorMean(diffy)
	meanDiffx, _ := ComputeVectorMean(diffx)
	fmt.Println("mean diff x", meanDiffx)
	fmt.Println("mean diff y", meanDiffy)
	//The -1 part below we should make input
	fmt.Println(meanDiffx * meanDiffy)
	coVar := meanDiffx * meanDiffy / (float64(len(x)) - 1)
	return coVar, nil

}

//Regressor fits LinReg struct on input values x and output y
func Regressor(x, y []float64) LinReg {
	//takes in x, y returns fitted LinReg

	// Coefs
	// B1 = sum((x(i) - mean(x)) * (y(i) - mean(y))) / sum( (x(i) - mean(x))^2 )
	// B0 = mean(y) - B1 * mean(x)
	// B1 = covariance(x, y) / variance(x)

	//check weights for fit
	return LinReg{1, []float64{1, 2, 3}}

}
