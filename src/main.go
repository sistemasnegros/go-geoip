package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/oschwald/geoip2-golang"
)

func main() {
	home, err := os.UserHomeDir()

	if err != nil {
		home = "./"
	}

	db, err := geoip2.Open(fmt.Sprintf("%s/GeoLite2-Country.mmdb", home))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	argsWithoutProg := os.Args[1:]

	for _, v := range argsWithoutProg {

		ip := net.ParseIP(v)

		if ip == nil {
			fmt.Printf("%s => err ip invalid\n", v)
			continue
		}

		record, err := db.Country(ip)
		if err != nil {
			fmt.Printf("%s => not found\n", v)
			continue
		}

		fmt.Printf("%s => %s \n", record.Country.Names["en"], v)

	}

}
