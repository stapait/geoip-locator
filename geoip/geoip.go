package geoip

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
)

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

	fmt.Println(inet.IP.String())
	fmt.Printf("CIDR: %s total: %d first: %s last: %s ones: %d\n", cidr, totalIPS, IntToIpv4(firstDecimalIP), IntToIpv4(lastDecimalIP), ones)

	return firstDecimalIP, lastDecimalIP
}
