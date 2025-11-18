package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TimeTestSuite struct {
	suite.Suite
}

func (s *TimeTestSuite) TestConvertNanoToString() {
	s.Assert().Equal("3m17.331667813s", ConvertNanoToString(197331667813.80))
	s.Assert().Equal("3m29.226522788s", ConvertNanoToString(209226522788))
	s.Assert().Equal("3m7.893379134s", ConvertNanoToString(187893379134.47))
	s.Assert().Equal("3m19.085533836s", ConvertNanoToString(199085533836.05))
	s.Assert().Equal("3m21.858747766s", ConvertNanoToString(201858747766.39))

	s.Assert().Equal("0.000000000s", ConvertNanoToString(0))
	s.Assert().Equal("3h6m41.858747766s", ConvertNanoToString(11201858747766.39))
}

func (s *TimeTestSuite) TestGetMinMaxTimes() {
	middle := time.Now()
	before := middle.Add(-1 * time.Hour)
	after := middle.Add(time.Hour)

	min, max := GetMinMaxTimes(&before, &after, nil)
	s.Assert().Equal(*min, before)
	s.Assert().Equal(*max, after)

	min, max = GetMinMaxTimes(nil, nil, &middle)
	s.Assert().Equal(*min, middle)
	s.Assert().Nil(max)

	min, max = GetMinMaxTimes(&middle, nil, &before)
	s.Assert().Equal(*min, before)
	s.Assert().Equal(*max, middle)

	min, max = GetMinMaxTimes(&before, nil, &middle)
	s.Assert().Equal(*min, before)
	s.Assert().Equal(*max, middle)

	min, max = GetMinMaxTimes(&before, nil, &before)
	s.Assert().Equal(*min, before)
	s.Assert().Nil(max)

	min, max = GetMinMaxTimes(&before, &middle, &after)
	s.Assert().Equal(*min, before)
	s.Assert().Equal(*max, after)

	min, max = GetMinMaxTimes(&middle, &after, &before)
	s.Assert().Equal(*min, before)
	s.Assert().Equal(*max, after)

	min, max = GetMinMaxTimes(&before, &after, &after)
	s.Assert().Equal(*min, before)
	s.Assert().Equal(*max, after)

}

func TestTimeTestSuite(t *testing.T) {
	suite.Run(t, new(TimeTestSuite))
}
