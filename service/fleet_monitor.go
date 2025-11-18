package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/codedepot/fleet-monitor/client"
	"github.com/codedepot/fleet-monitor/datastore"
	"github.com/codedepot/fleet-monitor/util"
	"github.com/gorilla/mux"
)

type FleetMonitorService struct {
	router *mux.Router

	store datastore.Datastore
}

func NewFleetMonitorService(router *mux.Router, store datastore.Datastore) *FleetMonitorService {
	fms := &FleetMonitorService{router: router, store: store}
	fms.router.Use(fms.loggingMiddleware)
	fms.setupDeviceEndpoints()

	return fms
}

func (fms *FleetMonitorService) setupDeviceEndpoints() {
	deviceRouter := fms.router.PathPrefix("/api/v1/devices/{device_id}").Subrouter()

	deviceRouter.HandleFunc("/heartbeat", fms.heartbeatHandler).Methods(http.MethodPost)

	deviceRouter.HandleFunc("/stats", fms.getStatsHandler).Methods(http.MethodGet)
	deviceRouter.HandleFunc("/stats", fms.postStatsHandler).Methods(http.MethodPost)
}

func (fms *FleetMonitorService) heartbeatHandler(w http.ResponseWriter, r *http.Request) {
	deviceId := mux.Vars(r)["device_id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, "no request body", http.StatusBadRequest)
		return
	}

	req := client.HeartbeatRequest{}
	err = json.Unmarshal(body, &req)

	if err != nil {
		writeError(w, "improper request body provided", http.StatusBadRequest)
		return
	}

	storeErr := fms.store.SaveHeartbeat(deviceId, req.GetSentAt())
	if storeErr != nil {
		writeStoreError(w, storeErr)
		return
	}

	writeResponse(w, "", http.StatusNoContent)

}

// create a manual request object for this endpoint as the one specified in the openapi spec does not match the test cases.
// The test cases use values that are bigger than int32 MAX_VALUE
type UploadStatsRequest struct {
	SentAt     time.Time `json:"sent_at"`
	UploadTime int64     `json:"upload_time"`
}

func (fms *FleetMonitorService) postStatsHandler(w http.ResponseWriter, r *http.Request) {
	deviceId := mux.Vars(r)["device_id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, "no request body", http.StatusBadRequest)
		return
	}
	req := UploadStatsRequest{}
	err = json.Unmarshal(body, &req)

	if err != nil {
		writeError(w, "improper request body provided", http.StatusBadRequest)
		return
	}
	storeErr := fms.store.SaveUploadTime(deviceId, req.SentAt, req.UploadTime)
	if storeErr != nil {
		writeStoreError(w, storeErr)
		return
	}

	writeResponse(w, "", http.StatusNoContent)
}

func (fms *FleetMonitorService) getStatsHandler(w http.ResponseWriter, r *http.Request) {
	deviceId := mux.Vars(r)["device_id"]

	stats, storeErr := fms.store.GetStats(deviceId)
	if storeErr != nil {
		writeStoreError(w, storeErr)
		return
	}

	writeJsonResponse(w, client.GetDeviceStatsResponse{
		AvgUploadTime: util.ConvertNanoToString(stats.AvgUploadTime),
		Uptime:        stats.Uptime,
	}, http.StatusOK)
}

func (fms *FleetMonitorService) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func writeStoreError(w http.ResponseWriter, err datastore.StoreError) {
	_, isNotFound := err.(*datastore.NotFoundError)
	statusCode := http.StatusInternalServerError
	if isNotFound {
		statusCode = http.StatusNotFound
	}

	writeError(w, err.Msg(), statusCode)
}

func writeError(w http.ResponseWriter, message string, statusCode int) {
	response := client.ErrorResponse{Msg: message}
	resStr, err := json.Marshal(response)
	if err != nil {
		fmt.Println("could not marshall error: ", err.Error())
	}
	writeResponse(w, string(resStr), statusCode)
}

func writeJsonResponse(w http.ResponseWriter, response any, statusCode int) {
	resBytes, err := json.Marshal(response)
	if err != nil {
		writeError(w, "could not marshall response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	if response != "" {
		w.Write(resBytes)
	}
}

func writeResponse(w http.ResponseWriter, response string, statusCode int) {
	w.WriteHeader(statusCode)
	if response != "" {
		w.Write([]byte(response))
	}
}
