package main

import (
	"flag"
	"fmt"
	"log"
	"ssctool/csvmapper"
)

var csv = flag.String("csvfile", "", "CSV file to read")
var field = flag.String("field", "", "Field to extract")

func main() {
	flag.Parse()

	if *csv == "" {
		log.Fatal("need -csvfile")
	}

	if *field == "" {
		log.Fatal("need -field")
	}

	data, err := csvmapper.ReadCSVWithHeaders(*csv)
	if err != nil {
		log.Fatalf("Failed to read CSV: %v", err)
	}

	for _, r := range data {
		if r[*field] != "" {
			fmt.Println("\"" + r[*field] + "\"")
		}
	}
}
