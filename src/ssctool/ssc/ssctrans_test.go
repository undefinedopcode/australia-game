package ssc_test

import (
	"ssctool/ssc"
	"testing"
)

func Test_Lookup(t *testing.T) {
	fn := "../../../data/geog.csv"
	s, err := ssc.New(fn)
	if err != nil {
		t.Fatalf("Error creating ssc translator: %v", err)
	}

	code := s.LookupCodeFromSuburb("Elmore")
	t.Logf("Code lookup gave: %s", code)
	if code != "SSC20455" {
		t.Fatalf("Expected Elmore to yield SSC20455, but got [%s]", code)
	}

}
