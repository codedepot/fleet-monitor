package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/codedepot/fleet-monitor/datastore"
	"github.com/codedepot/fleet-monitor/service"
	"github.com/codedepot/fleet-monitor/util"
	"github.com/gorilla/mux"
)

func main() {
	host := util.GetOptionalStringVariable("HOST", "127.0.0.1")
	port := util.GetOptionalStringVariable("PORT", "6733")
	csvLocation := util.GetOptionalStringVariable("DEVICES_PATH", "./devices.csv")

	router := mux.NewRouter()

	csvData, err := util.ReadCsv(csvLocation)
	if err != nil {
		log.Fatal(fmt.Sprintf("could not read from csv file in location %s. Error: %s", csvLocation, err.Error()))
	}
	// skip first row
	deviceIds := util.GetColumnData(0, csvData[1:])

	store := datastore.NewInMemoryDatastore()
	store.Register(deviceIds)

	service.NewFleetMonitorService(router, store)
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%s:%s", host, port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
