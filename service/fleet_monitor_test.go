package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/codedepot/fleet-monitor/client"
	"github.com/codedepot/fleet-monitor/datastore"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

type FleetMonitorSuite struct {
	suite.Suite
	server *httptest.Server
}

func (s *FleetMonitorSuite) SetupTest() {

	store := datastore.NewInMemoryDatastore()
	store.Register([]string{"0", "1", "2"})
	router := mux.NewRouter()

	NewFleetMonitorService(router, store)

	s.server = httptest.NewServer(router)

}

func (s *FleetMonitorSuite) TeardownTest() {
	s.server.Close()

}

func (s *FleetMonitorSuite) TestHeartbeatHandler() {
	req := client.HeartbeatRequest{SentAt: time.Now()}
	jsonStr, err := json.Marshal(req)
	s.Assert().NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/api/v1/devices/1/heartbeat", s.server.URL), "application/json", bytes.NewBuffer(jsonStr))
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusNoContent, resp.StatusCode)
}

func (s *FleetMonitorSuite) TestHeartbeatHandlerError() {
	resp, err := http.Post(fmt.Sprintf("%s/api/v1/devices/1/heartbeat", s.server.URL), "application/json", http.NoBody)
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusBadRequest, resp.StatusCode)

	req := client.HeartbeatRequest{SentAt: time.Now()}
	jsonStr, err := json.Marshal(req)
	s.Assert().NoError(err)
	resp, err = http.Post(fmt.Sprintf("%s/api/v1/devices/5/heartbeat", s.server.URL), "application/json", bytes.NewBuffer(jsonStr))
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusNotFound, resp.StatusCode)
}

func (s *FleetMonitorSuite) TestPostStatsHandler() {
	req := UploadStatsRequest{SentAt: time.Now(), UploadTime: int64(24234)}
	jsonStr, err := json.Marshal(req)
	s.Assert().NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/api/v1/devices/1/stats", s.server.URL), "application/json", bytes.NewBuffer(jsonStr))
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusNoContent, resp.StatusCode)
}

func (s *FleetMonitorSuite) TestPostStatsHandlerError() {
	resp, err := http.Post(fmt.Sprintf("%s/api/v1/devices/1/stats", s.server.URL), "application/json", http.NoBody)
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusBadRequest, resp.StatusCode)

	req := UploadStatsRequest{SentAt: time.Now(), UploadTime: int64(24234)}
	jsonStr, err := json.Marshal(req)
	s.Assert().NoError(err)
	resp, err = http.Post(fmt.Sprintf("%s/api/v1/devices/5/stats", s.server.URL), "application/json", bytes.NewBuffer(jsonStr))
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusNotFound, resp.StatusCode)
}

func (s *FleetMonitorSuite) TestGetStatsHandler() {
	timeStr := "2025-11-10T13:30:00Z"
	time1, err := time.Parse(time.RFC3339, timeStr)
	time2 := time1.Add(time.Hour)

	heartbeatReq := client.HeartbeatRequest{SentAt: time1}
	jsonStr, err := json.Marshal(heartbeatReq)
	s.Assert().NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/api/v1/devices/1/heartbeat", s.server.URL), "application/json", bytes.NewBuffer(jsonStr))
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusNoContent, resp.StatusCode)

	heartbeatReq = client.HeartbeatRequest{SentAt: time2}
	jsonStr, err = json.Marshal(heartbeatReq)
	s.Assert().NoError(err)
	resp, err = http.Post(fmt.Sprintf("%s/api/v1/devices/1/heartbeat", s.server.URL), "application/json", bytes.NewBuffer(jsonStr))
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusNoContent, resp.StatusCode)

	statsReq := UploadStatsRequest{SentAt: time2, UploadTime: int64(1000000)}
	jsonStr, err = json.Marshal(statsReq)
	s.Assert().NoError(err)
	resp, err = http.Post(fmt.Sprintf("%s/api/v1/devices/1/stats", s.server.URL), "application/json", bytes.NewBuffer(jsonStr))
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusNoContent, resp.StatusCode)

	resp, err = http.Get(fmt.Sprintf("%s/api/v1/devices/1/stats", s.server.URL))
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusOK, resp.StatusCode)
	respStr, err := io.ReadAll(resp.Body)
	s.Assert().NoError(err)

	stats := client.GetDeviceStatsResponse{}
	err = json.Unmarshal(respStr, &stats)
	s.Assert().NoError(err)
	s.Assert().Equal(float64(3.3333333333333335), stats.Uptime)
	s.Assert().Equal("0.001000000s", stats.AvgUploadTime)
}

func (s *FleetMonitorSuite) TestGetStatsHandlerError() {
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/devices/5/stats", s.server.URL))
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusNotFound, resp.StatusCode)
}

func TestFleetMonitorSuite(t *testing.T) {
	suite.Run(t, new(FleetMonitorSuite))
}
