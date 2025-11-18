package datastore

import (
	"math/big"
	"sync"
	"time"

	"github.com/codedepot/fleet-monitor/util"
)

type Datastore interface {
	Register(deviceIds []string)
	SaveHeartbeat(deviceId string, timestamp time.Time) StoreError
	SaveUploadTime(deviceId string, timestamp time.Time, uploadTime int64) StoreError
	GetStats(deviceId string) (*DeviceStats, StoreError)
}

type DeviceStats struct {
	Uptime        float64
	AvgUploadTime float64
}

type DeviceData struct {
	heartbeatCount int64
	firstHeartbeat *time.Time
	lastHeartbeat  *time.Time

	uploadTimeDurationSum *big.Int
	uploadTimeCount       int64
}

type InMemoryDatastore struct {
	lock sync.RWMutex
	data map[string]DeviceData
}

func NewInMemoryDatastore() *InMemoryDatastore {
	return &InMemoryDatastore{lock: sync.RWMutex{}, data: make(map[string]DeviceData)}

}

func (imd *InMemoryDatastore) Register(deviceIds []string) {
	imd.lock.Lock()
	for _, id := range deviceIds {
		imd.data[id] = DeviceData{
			uploadTimeDurationSum: big.NewInt(0),
		}
	}
	imd.lock.Unlock()
}

func (imd *InMemoryDatastore) SaveHeartbeat(deviceId string, timestamp time.Time) StoreError {
	imd.lock.Lock()
	val, ok := imd.data[deviceId]

	if !ok {
		imd.lock.Unlock()
		return &NotFoundError{msg: "device not found"}
	}

	min, max := util.GetMinMaxTimes(val.firstHeartbeat, val.lastHeartbeat, &timestamp)
	val.firstHeartbeat = min
	val.lastHeartbeat = max

	val.heartbeatCount++
	imd.data[deviceId] = val
	imd.lock.Unlock()
	return nil
}

func (imd *InMemoryDatastore) SaveUploadTime(deviceId string, timestamp time.Time, uploadTime int64) StoreError {
	imd.lock.Lock()
	val, ok := imd.data[deviceId]

	if !ok {
		imd.lock.Unlock()
		return &NotFoundError{msg: "device not found"}
	}

	val.uploadTimeDurationSum = big.NewInt(0).Add(val.uploadTimeDurationSum, big.NewInt(uploadTime))
	val.uploadTimeCount++
	imd.data[deviceId] = val

	imd.lock.Unlock()
	return nil
}

func (imd *InMemoryDatastore) GetStats(deviceId string) (*DeviceStats, StoreError) {
	imd.lock.RLock()
	val, ok := imd.data[deviceId]
	imd.lock.RUnlock()
	if !ok {
		return nil, &NotFoundError{msg: "device not found"}
	}
	if val.firstHeartbeat == nil || val.lastHeartbeat == nil {
		return nil, &NotFoundError{msg: "not enough data, need at least two heartbeats"}
	}

	difference := val.lastHeartbeat.Sub(*val.firstHeartbeat)
	uptime := 100 * float64(val.heartbeatCount) / float64(difference.Minutes())
	sumFloat, _ := val.uploadTimeDurationSum.Float64()
	timeDurationAvg := sumFloat / float64(val.uploadTimeCount)

	return &DeviceStats{
		Uptime:        uptime,
		AvgUploadTime: timeDurationAvg,
	}, nil
}
