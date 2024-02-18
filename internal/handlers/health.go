package handlers

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	var statusMsg, statusCode = checkDatabaseConnection()
	s, err := json.Marshal(Status{
		Code: statusCode,
		Msg:  statusMsg,
	})
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(s)
}

func checkDatabaseConnection() (string, int) {
	return "OK", http.StatusOK
}
