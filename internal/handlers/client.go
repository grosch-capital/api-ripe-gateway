package handlers

import "net/http"

func RAWIpInformationHandler(w http.ResponseWriter, r *http.Request) {
	IPAddress := r.Header.Get("X-Real-Ip")
	w.Write([]byte(IPAddress))

	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
		w.Write([]byte(IPAddress))
	}

	if IPAddress == "" {
		IPAddress = r.RemoteAddr
		w.Write([]byte(IPAddress))
	}

}

func JSONIpInformationHandler(w http.ResponseWriter, r *http.Request) {
	ip_addr := r.Header.Get("X-Real-Ip")
	if ip_addr == "" {
		ip_addr = r.Header.Get("X-Forwarded-For")
	}
	if ip_addr == "" {
		ip_addr = r.RemoteAddr
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"ip_address":"` + ip_addr + `"}`))
}

// func RAWGeoInformationHandler(w http.ResponseWriter, r *http.Request) {

// }

// func JSONGeoInformationHandler(w http.ResponseWriter, r *http.Request) {

// }
