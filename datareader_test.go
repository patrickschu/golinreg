//tests for datareader
package regression

import (
	"fmt"
	"testing"
)

func TestReadCsv(t *testing.T) {
	sampleFile := "/Users/pschultz/Documents/golinreg/testframe.csv"
	df, err := ReadCsvRows(sampleFile)
	if err != nil {
		t.Errorf("Error %d raised, want: raise", err)
	}
	fmt.Println(df)

	//comparevectors()

}
