package regression

import (
	"fmt"
	"testing"
)

//try these tests

func TestGetErrors(t *testing.T) {
	//Error out if length different
	got, err := GetErrors([]float64{1, 2, 3}, []float64{100, 200})
	fmt.Println(got, err)
	if err == nil {
		t.Errorf("different input len, error: %s want: raised", err)
	}
	//Error out if length is 0
	got1, err1 := GetErrors([]float64{}, []float64{})
	fmt.Println(got1, err1)
	if err1 == nil {
		t.Errorf("zero input len, error: %s want: raised", err)
	}
	//test legit input; preds is first
	got2, err2 := GetErrors([]float64{1, 2, 3}, []float64{100.5, 200.5, 300.5})
	want2 := []float64{99.5, 198.5, 297.5}
	fmt.Println(got2, err2)
	if err2 != nil {
		t.Errorf("legit input, error: %s want: nil", err)
	}
	for ind, outputVal := range got2 {
		fmt.Println(outputVal - want2[ind])
		if outputVal != want2[ind] {
			t.Errorf("incorrect output, want: %e, got: %e", want2[ind], outputVal)
		}
	}
}
func TestVectorSquares(t *testing.T) {
	// test VectorSquares
	// This is not good since depends on GetErrors
	got2, _ := GetErrors([]float64{1, 2, 3}, []float64{100.5, 200.5, 300.5})
	want3 := []float64{9900.25, 39402.25, 88506.25}
	got3, err3 := SquareVector(got2)
	if err3 != nil {
		t.Errorf("error should be nil,  got: %s", err3)
	}
	for ind, outputVal := range got3 {
		fmt.Println(outputVal - want3[ind])
		if outputVal != want3[ind] {
			t.Errorf("incorrect output, want: %e, got: %e", want3[ind], outputVal)
		}
	}

}
