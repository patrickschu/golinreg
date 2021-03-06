//Put regression code here
package regression

import (
	"errors"
	"fmt"
	"math"
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

//GetErrors of `preds` vs `truth`
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

//ComputeGradient -- NOT USED
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

//Mean of vector x
func Mean(x []float64) (float64, error) {
	vectorsum, _ := SumVector(x)
	mean := vectorsum / float64(len(x))
	return mean, nil
}

//Variance returns variance for x; nCorrect is subtracted from denominator ie needs to be 1 to match R
func Variance(x []float64, nCorrect int) (float64, error) {
	diffs := make([]float64, len(x))
	meanx, _ := Mean(x)
	//subtract meanx from each item in x, this we can outsource
	for ind, number := range x {
		diff := number - meanx
		diffs[ind] = diff * diff
	}
	sum, _ := SumVector(diffs)
	variance := sum / float64(len(x)-nCorrect)
	return variance, nil
}

//Stdev returns standard deviation of x
func Stdev(x []float64) (float64, error) {
	variance, err := Variance(x, 1)
	return math.Sqrt(variance), err
}

//Correlation computes the correlation between to inputs
func Correlation(x []float64, y []float64) (float64, error) {
	stdevx, _ := Stdev(x)
	stdevy, _ := Stdev(y)
	if stdevx == 0 || stdevy == 0 {
		return 0.0, nil
	}
	covar, _ := Covariance(x, y)
	correlation := covar / stdevx / stdevy
	return correlation, nil
}

//AddtoVector adds number `add` to every item in x
func AddtoVector(x []float64, add float64) ([]float64, error) {
	outputVector := make([]float64, len(x))
	for ind, number := range x {
		result := number + add
		fmt.Println("number", number, "out", result)
		outputVector[ind] = result
	}
	return outputVector, nil
}

// Covariance between x and y (uses n-1)
func Covariance(x, y []float64) (float64, error) {
	if len(x) != len(y) {
		return 0, fmt.Errorf(
			"Covariance: input vectors are len %d , %d, needs to be ==", len(x), len(y))
	}
	meanx, _ := Mean(x)
	meany, _ := Mean(y)
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
func Regressor(x []float64, y []float64) (LinReg, error) {
	//takes in x, y returns fitted LinReg
	meany, _ := Mean(y)
	meanx, _ := Mean(x)
	var intercept float64
	//correlation (x,y) * stdev(y) / stdev(x)
	//weight := (covar / variance)
	correlation, err := Correlation(x, y)
	if err != nil {
		fmt.Errorf("Regressor: Correlation returns %s", err)
		return LinReg{}, err
	}
	stdevx, _ := Stdev(x)
	stdevy, _ := Stdev(y)
	weight := correlation * (stdevy / stdevx)
	intercept = meany - (weight * meanx)
	fmt.Printf("weight is %e", weight)
	fmt.Printf("intercept is %e", intercept)
	return LinReg{intercept, []float64{weight}}, nil
}

// Coefs
// B1 = sum((x(i) - mean(x)) * (y(i) - mean(y))) / sum( (x(i) - mean(x))^2 )
// B0 = mean(y) - B1 * mean(x)
// B1 = covariance(x, y) / variance(x)
