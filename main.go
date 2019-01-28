package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

const countriesCSVFile = "files/GeoLite2-Country-Locations-en.csv"
const countryBlocksCSVFile = "GeoLite2-Country-Blocks-IPv4"

type Country struct {
	ContinentCode string
	ContinentName string
	CountryCode   string
	CountryName   string
}

type IPRange struct {
	RangeStart float64
	RangeEnd   float64
	CountryID  int
}

var countries map[int]Country
var ipRanges []IPRange

func loadCountriesFromCSV(filename string) {
	csvFile, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error opning file %s: %+v", filename, err)
		os.Exit(1)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	defer csvFile.Close()

	reader.Read() // Skip field names

	countries = make(map[int]Country)

	for {
		s, err := reader.Read()

		if err == io.EOF {
			break
		}
		countryCode, _ := strconv.Atoi(s[0])

		countries[countryCode] = Country{
			ContinentCode: s[2],
			ContinentName: s[3],
			CountryCode:   s[4],
			CountryName:   s[5],
		}
	}
}

func loadIPRangesFromCSV(fileName string) {

}

func main() {
	loadCountriesFromCSV(countriesCSVFile)

	loadIPRangesFromCSV(countryBlocksCSVFile)
}
