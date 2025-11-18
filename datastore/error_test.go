package datastore

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ErrorTestSuite struct {
	suite.Suite
}

func (s *ErrorTestSuite) TestNotFoundError() {
	err := NotFoundError{msg: "my-message"}
	s.Assert().Equal("my-message", err.Msg())
}

func (s *ErrorTestSuite) TestServerError() {
	err := ServerError{msg: "my-message"}
	s.Assert().Equal("my-message", err.Msg())
}

func TestErrorTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorTestSuite))
}
