package ssc

import (
	"log"
	"ssctool/csvmapper"
	"strings"
)

/*
SSC manages a CSV mapping of SSCCODE to Suburb and vice-versa
*/

// SSC represents a valid id prefix, as the mapping file may contain other
// types
const SSCPrefix = "SSC"

type SSCTranslator struct {
	ssc2suburbname map[string]string
	suburbname2ssc map[string]string
	Filename       string
}

func New(filename string) (*SSCTranslator, error) {
	s := &SSCTranslator{
		Filename:       filename,
		ssc2suburbname: make(map[string]string),
		suburbname2ssc: make(map[string]string),
	}
	err := s.ReadData()
	return s, err
}

func (s *SSCTranslator) ReadData() error {

	m, err := csvmapper.ReadCSVWithHeaders(s.Filename)
	if err != nil {
		return err
	}

	for _, r := range m {
		if code, ok := r["Code"]; ok {
			if label, ok := r["Label"]; ok {
				if !strings.HasPrefix(code, SSCPrefix) {
					continue
				}
				s.ssc2suburbname[code] = label
				s.suburbname2ssc[strings.ToLower(label)] = code
				log.Printf("SSCTranslator: loaded mapping %s -> %s", code, label)
			}
		}
	}

	return nil

}

func (s *SSCTranslator) LookupCodeFromSuburb(suburb string) string {
	return s.suburbname2ssc[strings.ToLower(suburb)]
}

func (s *SSCTranslator) LookupSuburbFromCode(code string) string {
	return s.ssc2suburbname[strings.ToUpper(code)]
}
