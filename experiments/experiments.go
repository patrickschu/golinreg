//
package experiments

import (
	"errors"
	"fmt"
)

// Compute dot product
func DotProd(vec1 []float64, vec2 []float64) (float64, error) {
	if len(vec1) != len(vec2) {
		return 0, errors.New(fmt.Sprintf("Input vectors to DotProd are not equal length: %d and %d", len(vec1), len(vec2)))
	}
	var result float64
	for ind, number := range vec1 {
		result += number * vec2[ind]

	}
	return result, nil

}



func DotProdVector(vec1 []float64, longvec [][]float64) ([]float64, error) {
	if len(vec1) != len(longvec[0]) {
		return []float64{}, errors.New(fmt.Sprintf("Input vectors to DotProd are not equal length: %d and %d", len(vec1), len(longvec[0])))
	}
	fmt.Println(vec1, longvec)
	// check if all items in longvec are same length
	result := make([]float64, len(longvec))
	for ind, vector := range longvec {
		fmt.Println(ind, vector, vec1)
		res, _ := DotProd(vec1, vector)
		result[ind] = res
	}
	if len(result) != len(longvec) {
		return []float64{}, errors.New(fmt.Sprintf("Result vector length != input vector length (%d != %d)", len(result), len(longvec)))
			}
	return result, nil
	}