package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EnvTestSuite struct {
	suite.Suite
}

func (s *EnvTestSuite) TestGetOptionalStringVariableSet() {
	os.Setenv("NAME", "Value")

	val := GetOptionalStringVariable("NAME", "Default")
	s.Assert().Equal("Value", val)
}

func (s *EnvTestSuite) TestGetOptionalIntVariableDefault() {
	os.Setenv("NAME", "")

	val := GetOptionalStringVariable("NAME", "Default")
	s.Assert().Equal("Default", val)
}

func TestEnvTestSuite(t *testing.T) {
	suite.Run(t, new(EnvTestSuite))
}
