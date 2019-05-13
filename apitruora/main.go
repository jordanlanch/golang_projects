package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jordanlanch/proyect/projects/apitruora/persistence"
)

func main() {
	// DB connection
	db, err := persistence.MigrateDB()
	if err != nil {
		log.Fatal(err)
	}
	// Router
	router := mux.NewRouter()

	router.HandleFunc("/server/{domain}", GetServerByDomain).Methods("GET")

	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// GetCassandraTrackerRegisterByPlate endpoint to read data filter by plate
func GetServerByDomain(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	plate := params["domain"]
	trackerData, err := storecassandra.GetTrackingByPlate(plate)
	if err != nil {
		sendInternalServerError(err, w)
	}
	jsonData, err := json.Marshal(trackerData)
	if err != nil {
		sendInternalServerError(err, w)
	}
	sendOkResponse(jsonData, w)

}