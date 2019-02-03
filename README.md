# Golang GeoIP locator

Retrieve country data given an IPv4 address. We only need a free GeoIP database, no paid dabases required.

Project uses Maxmind GeoLite2 Free Database in CSV format.

## Usage

First of all, download current zip file from https://dev.maxmind.com/geoip/geoip2/geolite2/, extract it and move these two CSV files into `files` folder:
* `GeoLite2-Country-Blocks-IPv4.csv`
* `GeoLite2-Country-Locations-en.csv`

With these two files we are able to locate any IP and get respective country.
