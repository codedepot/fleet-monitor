package datastore

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type DatastoreTestSuite struct {
	suite.Suite
	store *InMemoryDatastore
}

func (s *DatastoreTestSuite) SetupTest() {
	s.store = NewInMemoryDatastore()
	s.store.Register([]string{"0", "1", "2"})
}

func (s *DatastoreTestSuite) TestRegister() {
	s.Assert().Equal(3, len(s.store.data))
	s.Assert().Equal(int64(0), s.store.data["0"].heartbeatCount)
	s.Assert().Nil(s.store.data["0"].firstHeartbeat)
	s.Assert().Nil(s.store.data["0"].lastHeartbeat)

	s.Assert().Equal(int64(0), s.store.data["0"].uploadTimeDurationSum.Int64())
	s.Assert().Equal(int64(0), s.store.data["0"].uploadTimeCount)
}

func (s *DatastoreTestSuite) TestSaveFirstHeartbeat() {
	now := time.Now()

	// sets first heartbeat
	s.store.SaveHeartbeat("0", now)
	s.Assert().Equal(now, *s.store.data["0"].firstHeartbeat)
	s.Assert().Nil(s.store.data["0"].lastHeartbeat)
	s.Assert().Equal(int64(1), s.store.data["0"].heartbeatCount)

	// replace first heartbeat and set last heartbeat
	oneHourAgo := now.Add(-1 * time.Hour)
	s.store.SaveHeartbeat("0", oneHourAgo)
	s.Assert().Equal(oneHourAgo, *s.store.data["0"].firstHeartbeat)
	s.Assert().Equal(now, *s.store.data["0"].lastHeartbeat)
	s.Assert().Equal(int64(2), s.store.data["0"].heartbeatCount)

	// replace first heartbeat, last one unchanged
	twoHoursAgo := now.Add(-2 * time.Hour)
	s.store.SaveHeartbeat("0", twoHoursAgo)
	s.Assert().Equal(twoHoursAgo, *s.store.data["0"].firstHeartbeat)
	s.Assert().Equal(now, *s.store.data["0"].lastHeartbeat)
	s.Assert().Equal(int64(3), s.store.data["0"].heartbeatCount)

	// a heartbeat that does not change first or last
	halfAnHourAgo := now.Add(-30 * time.Minute)
	s.store.SaveHeartbeat("0", halfAnHourAgo)
	s.Assert().Equal(twoHoursAgo, *s.store.data["0"].firstHeartbeat)
	s.Assert().Equal(now, *s.store.data["0"].lastHeartbeat)
	s.Assert().Equal(int64(4), s.store.data["0"].heartbeatCount)

	anHourFromNow := now.Add(time.Hour)
	s.store.SaveHeartbeat("0", anHourFromNow)
	s.Assert().Equal(twoHoursAgo, *s.store.data["0"].firstHeartbeat)
	s.Assert().Equal(anHourFromNow, *s.store.data["0"].lastHeartbeat)
	s.Assert().Equal(int64(5), s.store.data["0"].heartbeatCount)
}

func (s *DatastoreTestSuite) TestSaveFirstHeartbeatNotFound() {
	err := s.store.SaveHeartbeat("5", time.Now())
	s.Assert().NotNil(err)
	s.Assert().Equal("device not found", err.Msg())
}

func (s *DatastoreTestSuite) TestSaveUploadTime() {
	err := s.store.SaveUploadTime("0", time.Now(), 50)
	s.Assert().Nil(err)
	s.Assert().Equal(int64(50), s.store.data["0"].uploadTimeDurationSum.Int64())
	s.Assert().Equal(int64(1), s.store.data["0"].uploadTimeCount)

	err = s.store.SaveUploadTime("0", time.Now(), 150)
	s.Assert().Nil(err)
	s.Assert().Equal(int64(200), s.store.data["0"].uploadTimeDurationSum.Int64())
	s.Assert().Equal(int64(2), s.store.data["0"].uploadTimeCount)

	err = s.store.SaveUploadTime("0", time.Now(), 300)
	s.Assert().Nil(err)
	s.Assert().Equal(int64(500), s.store.data["0"].uploadTimeDurationSum.Int64())
	s.Assert().Equal(int64(3), s.store.data["0"].uploadTimeCount)

}

func (s *DatastoreTestSuite) TestSaveUploadTimeNotFound() {
	err := s.store.SaveUploadTime("5", time.Now(), 0)
	s.Assert().NotNil(err)
	s.Assert().Equal("device not found", err.Msg())
}

func (s *DatastoreTestSuite) TestGetStats() {
	now := time.Now()
	oneHourAgo := now.Add(-1 * time.Hour)
	anHourFromNow := now.Add(time.Hour)

	s.store.SaveHeartbeat("0", oneHourAgo)
	s.store.SaveHeartbeat("0", now)
	s.store.SaveHeartbeat("0", anHourFromNow)

	s.store.SaveUploadTime("0", time.Now(), 200)
	s.store.SaveUploadTime("0", time.Now(), 400)

	stats, err := s.store.GetStats("0")
	s.Assert().Nil(err)
	s.Assert().Equal(float64(2.5), stats.Uptime)
	s.Assert().Equal(float64(300), stats.AvgUploadTime)

}

func (s *DatastoreTestSuite) TestGetStatsError() {
	_, err := s.store.GetStats("5")
	s.Assert().NotNil(err)
	s.Assert().Equal("device not found", err.Msg())

	_, err = s.store.GetStats("0")
	s.Assert().NotNil(err)
	s.Assert().Equal("not enough data, need at least two heartbeats", err.Msg())
}

// This test ensures that if many operations are done in parallel, there is no read/write conflict
func (s *DatastoreTestSuite) TestParallization() {
	waitGroup := sync.WaitGroup{}

	now := time.Now()
	anHourAgo := now.Add(-1 * time.Hour)
	twoHoursAgo := now.Add(-2 * time.Hour)
	anHourFromNow := now.Add(time.Hour)
	twoHoursFromNow := now.Add(2 * time.Hour)

	cases := []time.Time{
		now,
		anHourAgo,
		now,
		twoHoursAgo,
		twoHoursAgo,
		anHourAgo,
		anHourFromNow,
		anHourFromNow,
		now,
		twoHoursFromNow,
	}

	for _, c := range cases {
		waitGroup.Add(1)

		go func() {
			s.store.SaveHeartbeat("0", c)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()

	s.Assert().Equal(int64(10), s.store.data["0"].heartbeatCount)
	s.Assert().Equal(twoHoursAgo, *s.store.data["0"].firstHeartbeat)
	s.Assert().Equal(twoHoursFromNow, *s.store.data["0"].lastHeartbeat)
}

func TestDatastoreTestSuite(t *testing.T) {
	suite.Run(t, new(DatastoreTestSuite))
}
