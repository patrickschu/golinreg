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

func TestSumVector(t *testing.T) {
	in := []float64{1.2, 23.23, 30000}
	want := 30024.43
	res, err := SumVector(in)
	if err != nil {
		t.Errorf("error should be nil,  got: %s", err)
	}
	if res != want {
		t.Errorf("incorrect output, want: %e, got: %e", want, res)
	}
	fmt.Println("TestComputeVectorSum completed")

}

func TestComputeVectorMean(t *testing.T) {
	in := []float64{1.2, 23.23, 30000}
	want := 30024.43 / 3
	res, err := ComputeVectorMean(in)
	if err != nil {
		t.Errorf("error should be nil,  got: %s", err)
	}
	if res != want {
		t.Errorf("incorrect output, want: %e, got: %e", want, res)
	}
	fmt.Println("TestComputeVectorMean completed")

}

func Testvariance(t *testing.T) {
	return
}

func TestComputeVectorVariance(t *testing.T) {
	in := []float64{1.2, 23.23, 30000}
	want := 299755871.0
	res, err := ComputeVectorVariance(in)
	if err != nil {
		t.Errorf("error should be nil,  got: %s", err)
	}
	if res != want {
		t.Errorf("incorrect output, want: %e, got: %e", want, res)
	}
	fmt.Println("TestComputeVectorVariance completed")

}

func TestCovariance(t *testing.T) {
	in1 := []float64{1, 2, 3.3}
	in2 := []float64{1, 2323, 22.1}
	// Look this up in book
	want := -103.4400 //-68.9600
	res, err := covariance(in1, in2)
	if err != nil {
		t.Errorf("error wanted : nil, raised: %e ", err)
	}
	if res != want {
		t.Errorf("incorrect output, want: %e, got: %e", want, res)
	}
}
