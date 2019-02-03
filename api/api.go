package api

import (
	"encoding/json"
	"geoip-locator/geoip"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/search", handleSearch)
}

func handleSearch(w http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ip := request.URL.Query().Get("ip")
	if ip == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	country, err := geoip.IPToCountry(ip)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	c, err := json.Marshal(country)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(c)
}
