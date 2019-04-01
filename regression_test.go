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
	feat0 := []float64{1.2, 2.4, 3}
	want0 := 0.8400000000000001
	feat1 := []float64{12, 3, 44}
	want1 := 464.33333333333337
	feat2 := []float64{24, 7, 88}
	want2 := 1824.3333333333333
	ins := [][]float64{feat0, feat1, feat2}
	wants := []float64{want0, want1, want2}
	for ind, input := range ins {
		res, err := ComputeVectorVariance(input)
		if err != nil {
			t.Errorf("error should be nil,  got: %s", err)
		}
		if res != wants[ind] {
			fmt.Println(res)
			fmt.Println(wants[ind])
			t.Errorf("incorrect output, want: %e, got: %e", wants[ind], res)
		}

	}
	fmt.Println("TestComputeVectorVariance completed")
}

//this to compare two vectors of numbers for ID
func comparevectors(x, y []float64) error {
	for ind, number := range x {
		if number != y[ind] {
			return fmt.Errorf("compare vectors, mismatch at ind %d : %f != %f", ind, number, y[ind])
		}
	}
	return nil
}

func TestAddtoVector(t *testing.T) {
	have := []float64{1, 2.2, -3}
	want := []float64{0, 1.2, -4}
	res, _ := AddtoVector(have, -1)
	//if the below is not Nil we want to break th
	comparevectors(res, want)

}

func TestComputeVectorCovariance(t *testing.T) {
	feat1 := []float64{12, 3, 44}
	feat2 := []float64{24, 7, 88}
	// 920.3333
	//want := 920.3333
	res, _ := ComputeVectorCovariance(feat2, feat1)
	fmt.Println("covar res", res)

}
