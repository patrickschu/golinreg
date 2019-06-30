//tests
package experiments

import (
	"fmt"
	"testing"
)

type arraytest struct {
	want    []float64
	have    []float64
	success bool
}

type numbertest struct {
	want    float64
	have    float64
	success bool
}

/*
>>> t = [1,2,3]
>>> tt = [100, 100, 100]
>>> numpy.dot(t, tt)
600
*/

func assert_result(testresult numbertest, t *testing.T) numbertest {
	//
	if testresult.have != testresult.want {
		t.Errorf("Result mismatch: want %v, got  %v", testresult.want, testresult.have)
	}
	testresult.success = true
	return testresult
}

func TestDotProd(t *testing.T) {
	vec1 := []float64{1, 2, 3}
	vec2 := []float64{100, 200, 300}
	dot1 := numbertest{}
	dot1.want = 1400
	res, err := DotProd(vec1, vec2)
	// check for error
	if err != nil {
		t.Errorf("DotProd returned error %d", err)
	}

	// check result
	dot1.have = res
	assert_result(dot1, t)

}


func TestDotVectors(t *testing.T) {
	vec1 := []float64{1, 2, 3}
	longvec := [][]float64{[]float64{100, 200, 300}, []float64{2,2,2}}
	dot2 := arraytest{}
	dot2.want = []float64{1400, 12}
	res, err := DotProdVector(vec1, longvec)
	if err != nil {
		t.Errorf("DotProd returned error %d", err)
	}
	dot2.have = res
	fmt.Println(dot2.have, dot2.want)
	//for ind, item := range dot2.have {
	//	if item != dot2.want[ind] {
	//	t.Errorf("Result mismatch: want %v, got  %v", dot2.want, dot2.have)
	//	}
	}



