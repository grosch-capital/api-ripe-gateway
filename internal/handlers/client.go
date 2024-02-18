package handlers

import "net/http"

func RAWIpInformationHandler(w http.ResponseWriter, r *http.Request) {
	IPAddress := r.Header.Get("X-Real-Ip")

	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
		w.Write([]byte(IPAddress))
	}

	if IPAddress == "" {
		IPAddress = r.RemoteAddr
		w.Write([]byte(IPAddress))
	}

	w.Write([]byte(IPAddress))
}

// func JSONIpInformationHandler(w http.ResponseWriter, r *http.Request) {

// }

// func RAWGeoInformationHandler(w http.ResponseWriter, r *http.Request) {

// }

// func JSONGeoInformationHandler(w http.ResponseWriter, r *http.Request) {

// }
