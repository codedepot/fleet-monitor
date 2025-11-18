package util

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CsvTestSuite struct {
	suite.Suite
}

// instead of making a file for testing purposes, will just use the devices.csv already in the repo
// alternatively, a dummy csv file could be made and stored in the util package just for testing purposes
func (s *CsvTestSuite) TestReadCsv() {
	data, err := ReadCsv("../devices.csv")
	s.Assert().NoError(err)
	s.Assert().NotNil(data)
}

func (s *CsvTestSuite) TestReadError() {
	_, err := ReadCsv("bad path")
	s.Assert().Error(err)
}

func (s *CsvTestSuite) TestGetColumnData() {

	columnData := GetColumnData(0, [][]string{
		{"0", "1", "2"},
		{"00", "11", "22"},
		{"000", "111", "222"},
	})
	s.Assert().Equal([]string{"0", "00", "000"}, columnData)
}

func TestCsvTestSuite(t *testing.T) {
	suite.Run(t, new(CsvTestSuite))
}
