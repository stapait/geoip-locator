package geoip

import "testing"

func TestCidrToIPV4Range(t *testing.T) {
	cidr := []string{"223.165.128.0/17", "223.27.19.144/32", "221.128.81.4/30", "221.0.0.0/3", "221.130.0.0/15"}
	expectedFirstIP := []int64{Ipv4ToInt("223.165.128.0"), Ipv4ToInt("223.27.19.144"), Ipv4ToInt("221.128.81.4"), Ipv4ToInt("192.0.0.0"), Ipv4ToInt("221.130.0.0")}
	expectedLastIP := []int64{Ipv4ToInt("223.165.255.255"), Ipv4ToInt("223.27.19.144"), Ipv4ToInt("221.128.81.7"), Ipv4ToInt("223.255.255.255"), Ipv4ToInt("221.131.255.255")}

	for i, c := range cidr {
		firstIP, lastIP := CidrToIntIpv4(c)

		if expectedFirstIP[i] != firstIP {
			t.Fatalf("First IP %d: Expected %d got %d", i, expectedFirstIP[i], firstIP)
		}

		if expectedLastIP[i] != lastIP {
			t.Fatalf("Last IP %d: Expected %d got %d", i, expectedLastIP[i], lastIP)
		}
	}
}

func TestIpv4ToNumber(t *testing.T) {
	ip := "200.221.10.20"
	expected := int64(3369929236)
	numberIP := Ipv4ToInt(ip)

	if numberIP != expected {
		t.Fatalf("Expected %d got %d", expected, numberIP)
	}
}

func TestNumberToIpv4(t *testing.T) {
	ip := "200.221.10.20"
	number := int64(3369929236)
	result := IntToIpv4(number)

	if ip != result {
		t.Fatalf("Expected %s got %s", ip, result)
	}
}
