package main

import (
	"fmt"
	"geoip-locator/geoip"
	"os"
)

const countriesCSVFile = "files/GeoLite2-Country-Locations-en.csv"
const countryBlocksCSVFile = "files/GeoLite2-Country-Blocks-IPv4.csv"

func main() {
	geoip.LoadCountriesFromCSV(countriesCSVFile)
	geoip.LoadIPsFromCSV(countryBlocksCSVFile)

	ip := "222.223.222.222"
	country, err := geoip.IPToCountry(ip)

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(0)
	}

	fmt.Printf("%+v", country)
}
