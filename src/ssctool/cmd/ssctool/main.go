package main

import (
	"flag"
	"fmt"
	"log"
	"ssctool/csvmapper"
	"ssctool/ssc"
)

var sscmap = flag.String("ssc", "", "SSC suburb file (CSV)")
var sscdata = flag.String("data", "", "Data file (CSV)")
var sscid = flag.String("sscidfield", "region_id", "Region field id")
var extractfield = flag.String("extract", "", "Field to extract")

func main() {
	flag.Parse()

	if *sscmap == "" {
		log.Fatal("ssc not specified")
	}

	if *sscdata == "" {
		log.Fatal("data not specified")
	}

	if *extractfield == "" {
		log.Fatal("extract not specified")
	}

	s, err := ssc.New(*sscmap)
	if err != nil {
		log.Fatalf("Failed to load mapper: %v", err)
	}

	datafile, err := csvmapper.ReadCSVWithHeaders(*sscdata)
	if err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	for _, r := range datafile {
		code := r[*sscid]
		if code == "" {
			continue
		}
		f := r[*extractfield]
		if f == "" {
			continue
		}
		suburb := s.LookupSuburbFromCode(code)
		if suburb == "" {
			log.Printf("Failed to lookup suburb from %s", code)
			continue
		}
		fmt.Printf("\"%s\",%s\n", suburb, f)
	}
}
