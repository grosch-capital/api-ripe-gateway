package main

import (
	"github.com/gorilla/mux"
	"github.com/grosch-capital/api-ripe-gateway/internal/handlers"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HealthCheckHandler)
	r.HandleFunc("/user", handlers.TempHandler).Methods("POST")
	r.HandleFunc("/user/{ident}", handlers.TempHandler).Methods("GET")
	r.HandleFunc("/user/{ident}", handlers.TempHandler).Methods("UPDATE")
	r.HandleFunc("/user/{ident}", handlers.TempHandler).Methods("DELETE")
	r.HandleFunc("/sync", handlers.TempHandler).Methods("POST")
	r.HandleFunc("/sync/{ident}", handlers.TempHandler).Methods("GET")
	r.HandleFunc("/sync/{ident}", handlers.TempHandler).Methods("UPDATE")
	r.HandleFunc("/sync/{ident}", handlers.TempHandler).Methods("DELETE")
	r.HandleFunc("/wallet", handlers.TempHandler).Methods("GET")
	r.HandleFunc("/wallet/{ident}", handlers.TempHandler).Methods("POST")
	r.HandleFunc("/wallet/{ident}", handlers.TempHandler).Methods("POST")
	r.HandleFunc("/healthz", handlers.HealthCheckHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
