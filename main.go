package main

import (
	"fmt"
	"geoip-locator/api"
	"geoip-locator/geoip"
	"net/http"
)

const (
	countriesCSVFile     = "files/GeoLite2-Country-Locations-en.csv"
	countryBlocksCSVFile = "files/GeoLite2-Country-Blocks-IPv4.csv"
	serverPort           = 8080
)

func main() {
	geoip.LoadCountriesFromCSV(countriesCSVFile)
	geoip.LoadIPsFromCSV(countryBlocksCSVFile)

	api.LoadRoutes()

	fmt.Println("Server listening at port", serverPort)
	http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)
}
