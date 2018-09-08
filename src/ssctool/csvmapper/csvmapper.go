package csvmapper

import (
	"encoding/csv"
	"os"
)

func ReadCSVWithHeaders(filename string) ([]map[string]string, error) {

	var out = make([]map[string]string, 0)

	f, err := os.Open(filename)
	if err != nil {
		return out, err
	}

	r := csv.NewReader(f)

	rows, err := r.ReadAll()
	if err != nil {
		return out, err
	}

	header := rows[0]
	data := rows[1:]

	for _, d := range data {
		m := make(map[string]string)
		for fidx, fname := range header {
			if fidx < len(d) {
				m[fname] = d[fidx]
			}
		}
		out = append(out, m)
	}

	return out, nil

}
