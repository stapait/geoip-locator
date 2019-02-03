# Golang GeoIP locator HTTP Server

A simple HTTP server that retrieves country data given an IPv4 address. We only need a free GeoIP database, no paid dabases required.

Project uses Maxmind GeoLite2 Free Database in CSV format.

## Usage

First of all, download current zip file from https://dev.maxmind.com/geoip/geoip2/geolite2/, extract it and move these two CSV files into `files` folder:
* `GeoLite2-Country-Blocks-IPv4.csv`
* `GeoLite2-Country-Locations-en.csv`
With only these two files we are able to locate any IP and get respective country.

Run the server:

```
$ go run main.go
Server listening at port 8080
```

Search for an IPv4 address:
```
$ curl -i localhost:8080/search?ip=192.156.4.3
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 03 Feb 2019 22:03:13 GMT
Content-Length: 103

{"ContinentCode":"NA","ContinentName":"North America","CountryCode":"US","CountryName":"United States"}
```

