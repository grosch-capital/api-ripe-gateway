package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func LookupRAWHandler(w http.ResponseWriter, r *http.Request) {
	IPAddress := r.Header.Get("X-Real-Ip")

	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}

	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}

	w.Write([]byte(IPAddress))
}

func LookupJSONHandler(w http.ResponseWriter, r *http.Request) {
	IPAddress := r.Header.Get("X-Real-Ip")

	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}

	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"public_ip":"` + IPAddress + `"}`))
}

func LookupClientGeoHandler(w http.ResponseWriter, r *http.Request) {
	IPAddress := r.Header.Get("X-Real-Ip")

	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}

	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}

	reqest := "http://ip-api.com/json/" + IPAddress
	resp, err := http.Get(reqest)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		w.Write(body)
	} else {
		w.Write([]byte("Error"))
	}
}

func LookupSpecGeoHandler(w http.ResponseWriter, r *http.Request) {
	addr := mux.Vars(r)["ip"]

	reqest := "http://ip-api.com/json/" + addr
	resp, err := http.Get(reqest)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		w.Write(body)
	} else {
		w.Write([]byte("Error"))
	}
}
