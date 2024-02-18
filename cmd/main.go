package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grosch-capital/api-ripe-gateway/internal/handlers"
)

func main() {
	r := mux.NewRouter()

	http.Handle("/", r)
	r.HandleFunc("/raw/ip", handlers.RAWIpInformationHandler)
	r.HandleFunc("/raw/geo", handlers.RAWGeoInformationHandler)
	r.HandleFunc("/json/ip", handlers.JSONIpInformationHandler)
	//r.HandleFunc("/json/geo", handlers.JSONGeoInformationHandler)
	r.HandleFunc("/healthz", handlers.HealthCheckHandler)

	http.ListenAndServe(":8080", r)
}
