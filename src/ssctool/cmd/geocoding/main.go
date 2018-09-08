package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"googlemaps.github.io/maps"
)

// geocoding api key
const apiKey = "AIzaSyDZqGVwihFyh101nzeb6nLWe0h944xiT9U"

var region = flag.String("region", "", "Region to get geocoding for")
var stdin = flag.Bool("stdin", false, "read from stdin")

func main() {
	flag.Parse()

	if *region == "" && !*stdin {
		log.Fatal("Need to supply a region with -region flag. Eg. \"Elmore\"")
	}

	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error creating client: %s", err)
	}

	if !*stdin {
		geocodeRequest(c, *region)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			*region = scanner.Text()
			geocodeRequest(c, *region)
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func geocodeRequest(c *maps.Client, address string) {
	r := &maps.GeocodingRequest{
		Address: address,
		Region:  "au",
	}

	res, err := c.Geocode(context.Background(), r)
	if err != nil {
		log.Fatalf("Error geocoding: %v", err)
	}
	//pretty.Println(res[0])

	if len(res) == 0 {
		return
	}

	mm := make(map[string]string)
	for _, ai := range res[0].AddressComponents {
		name := ai.Types[0]
		ln := ai.LongName
		mm[name] = ln
	}

	fmt.Printf(
		"\"%s\",\"%s\",\"%s\",%s,%f,%f,%f,%f,%f,%f\n",
		address,
		mm["administrative_area_level_2"],
		mm["administrative_area_level_1"],
		mm["postal_code"],
		res[0].Geometry.Location.Lat,
		res[0].Geometry.Location.Lng,
		res[0].Geometry.Bounds.NorthEast.Lat,
		res[0].Geometry.Bounds.NorthEast.Lng,
		res[0].Geometry.Bounds.SouthWest.Lat,
		res[0].Geometry.Bounds.SouthWest.Lng,
	)
}
