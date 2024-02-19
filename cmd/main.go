package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grosch-capital/api-ripe-gateway/internal/handlers"
)

func main() {
	r := mux.NewRouter()

	x, _ := handlers.Traceroute("google.com")
	print(x)

	r.HandleFunc("/", handlers.LookupRAWHandler)
	r.HandleFunc("/ip/raw", handlers.LookupRAWHandler)
	r.HandleFunc("/ip/json", handlers.LookupJSONHandler)
	r.HandleFunc("/geo", handlers.LookupClientGeoHandler)
	r.HandleFunc("/geo/{ip}", handlers.LookupSpecGeoHandler)
	r.HandleFunc("/route", handlers.HealthCheckHandler)
	r.HandleFunc("/route/{host}", handlers.HealthCheckHandler)
	r.HandleFunc("/healthz", handlers.HealthCheckHandler)
	http.ListenAndServe(":8080", r)
}
