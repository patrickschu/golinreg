//tests for datareader
package csvreader

import (
	"fmt"
	"testing"
)

//run over two vectors to check for equality and such
func comparevectors(x, y []string) (bool, error) {
	for ind, val := range x {
		fmt.Println(val, y[ind])
		if val != y[ind] {
			return false, fmt.Errorf("vector mismatch at index %d: '%s' in x, '%s' in y", ind, val, y[ind])
		}
	}
	return true, nil
}

type testobject struct {
	want []string
	have []string
}

func TestReadCsvByRow(t *testing.T) {
	// check out the below whitespace, this is cray
	row1 := []string{"\ufefffeature1", "feature2", "3feature"}
	row2 := []string{"1", "apple", "None"}
	row3 := []string{"13", "bananas", "empty below"}
	row4 := []string{"4.5", "pear", ""}
	wanted := [][]string{row1, row2, row3, row4}

	sampleFile := "/Users/pschultz/Documents/golinreg/testframe.csv"
	df, err := ReadCsvRows(sampleFile)
	if err != nil {
		t.Errorf("Error %d raised, want: not raise", err)
	}
	for ind, val := range df {
		fmt.Printf("\n\nrun on %d , %q", ind, val)
		test := testobject{want: wanted[ind], have: val}
		res, erro := comparevectors(test.want, test.have)
		fmt.Printf("row matching returns %v", res)
		fmt.Printf("error is  %v", erro)
		if erro != nil {
			t.Error(erro)
		}
	}

}

func TestReadCsvByCol(t *testing.T) {
	// check out the below whitespace, this is cray
	/*
	   feature1	feature2	3feature
	   1	apple	None
	   13	bananas	empty below
	   4.5	pear
	*/

	col1 := []string{"\ufefffeature1", "1", "13", "4.5"}
	col2 := []string{"feature2", "apple", "bananas", "pear"}
	col3 := []string{"3feature", "None", "empty below"}
	wanted := [][]string{col1, col2, col3}

	sampleFile := "/Users/pschultz/Documents/golinreg/testframe.csv"
	df, err := ReadCsvCols(sampleFile)
	if err != nil {
		t.Errorf("Error %d raised, want: not raise", err)
	}
	for ind, val := range df {
		fmt.Printf("\n\nrun on %d , %q", ind, val)
		test := testobject{want: wanted[ind], have: val}
		res, erro := comparevectors(test.want, test.have)
		fmt.Printf("row matching returns %v", res)
		fmt.Printf("error is  %v", erro)
		if erro != nil {
			t.Error(erro)
		}
	}

}
