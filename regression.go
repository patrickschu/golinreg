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
	fmt.Println(outputDiffs)
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

func variance(xval float64, mean float64) float64 {
	// 	variance = sum( (x - mean(x))^2 )
	diff := xval - mean
	return diff * diff
}

//ComputeVectorVariance returns vector of variances for x
func ComputeVectorVariance(x []float64) ([]float64, error) {
	outputVariances := make([]float64, 0)
	meanx, _ := ComputeVectorMean(x)
	for _, number := range x {
		outputVariances = append(outputVariances, variance(number, meanx))
	}
	print("variances", outputVariances)
	return outputVariances, nil
}

func covariance(x, y []float64) (float64, error) {
	var covar float64
	// covariance = sum((x(i) - mean(x)) * (y(i) - mean(y)))
	meanx, _ := ComputeVectorMean(x)
	meany, _ := ComputeVectorMean(y)
	for ind, number := range x {
		xval := number - meanx
		yval := y[ind] - meany
		covar += (xval * yval)
	}
	fmt.Println("covar is", covar)
	return covar, nil
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
