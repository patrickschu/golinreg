//Data read functions
package csvreader

import (
	"encoding/csv"
	"io"
	"os"
)

//ReadCsvRows reads file from `inputFilePath` and stores in map by row
func ReadCsvRows(inputFilePath string) (map[int64][]string, error) {
	// open file, give to `NewReader`
	var ind int64
	dataframe := make(map[int64][]string)

	fileObject, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer fileObject.Close()
	csvReader := csv.NewReader(fileObject)
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		dataframe[ind] = row
		ind++
	}
	return dataframe, nil
}

//ReadCsvCols reads file from `inputFilePath` and stores in map by col
func ReadCsvCols(inputFilePath string) (map[int][]string, error) {
	dataframe := make(map[int][]string)
	fileObject, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer fileObject.Close()
	csvReader := csv.NewReader(fileObject)
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		for ind, item := range row {
			dataframe[ind] = append(dataframe[ind], item)
		}
	}
	return dataframe, nil
}
