package geoip

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
)

type Country struct {
	ContinentCode string
	ContinentName string
	CountryCode   string
	CountryName   string
}

type IPRange struct {
	RangeStart int64
	RangeEnd   int64
	CountryID  int
}

var countries map[int]Country
var ipRanges []IPRange

func Ipv4ToInt(ip string) int64 {
	var result int64
	ipGroups := strings.Split(ip, ".")

	for i := 0; i <= 3; i++ {
		value, _ := strconv.ParseFloat(ipGroups[i], 10)
		result += int64(value * math.Pow(256, float64(3-i)))
	}

	return result
}

func IntToIpv4(numberIP int64) string {
	b0 := strconv.FormatInt((numberIP>>24)&0xff, 10)
	b1 := strconv.FormatInt((numberIP>>16)&0xff, 10)
	b2 := strconv.FormatInt((numberIP>>8)&0xff, 10)
	b3 := strconv.FormatInt((numberIP & 0xff), 10)

	return fmt.Sprintf("%s.%s.%s.%s", b0, b1, b2, b3)
}

func CidrToIntIpv4(cidr string) (int64, int64) {
	_, inet, _ := net.ParseCIDR(cidr)
	ones, size := inet.Mask.Size()
	remainingBits := size - ones
	totalIPS := int64(math.Pow(2, float64(remainingBits)))

	firstDecimalIP := Ipv4ToInt(inet.IP.String())
	lastDecimalIP := firstDecimalIP

	if totalIPS > 1 {
		lastDecimalIP = firstDecimalIP + totalIPS - 1
	}

	return firstDecimalIP, lastDecimalIP
}

func LoadCountriesFromCSV(filename string) {
	csvFile, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error opning file %s: %+v", filename, err)
		os.Exit(1)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	defer csvFile.Close()

	reader.Read() // Skip field names

	countries = make(map[int]Country, 300)

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

func LoadIPsFromCSV(filename string) {
	csvFile, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error opning file %s: %+v", filename, err)
		os.Exit(1)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	defer csvFile.Close()

	reader.Read() // Skip field names

	ipRanges = make([]IPRange, 301000)

	for {
		s, err := reader.Read()

		if err == io.EOF {
			break
		}

		ipRangeStart, ipRangeEnd := CidrToIntIpv4(s[0])
		countryID, _ := strconv.ParseInt(s[1], 10, 32)
		ipRanges = append(ipRanges, IPRange{RangeStart: ipRangeStart, RangeEnd: ipRangeEnd, CountryID: int(countryID)})
	}
}

func IPToCountry(ip string) (Country, error) {
	intIP := Ipv4ToInt(ip)

	for _, i := range ipRanges {
		if intIP >= i.RangeStart && intIP <= i.RangeEnd {

			return countries[i.CountryID], nil
		}
	}

	return Country{}, fmt.Errorf("Country IP %s not found", ip)
}
