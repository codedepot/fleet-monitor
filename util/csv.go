package util

import (
	"encoding/csv"
	"os"
)

func ReadCsv(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetColumnData(colIndex int, csvData [][]string) []string {
	output := []string{}
	for _, row := range csvData {
		output = append(output, row[colIndex])
	}
	return output
}
