package medianrents

import (
	"fmt"
	"math/rand"
	"quiz/database"
	"quiz/question"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func New() (*question.Question, error) {

	q := "What was the median rent of %s in the %s in March, 2018?"
	r, err := database.Random4DistinctMedianRents()
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var an [4]question.Answer
	var locs [4]string
	var states [4]string
	var nelat [4]float64
	var nelong [4]float64
	var swlat [4]float64
	var swlong [4]float64
	var clat [4]float64
	var clon [4]float64
	idx := 0
	for r.Next() {
		var locality string
		var amount int
		var municipality string
		var state string
		var lat float64
		var long float64
		var bnelat float64
		var bnelong float64
		var bswlat float64
		var bswlong float64
		err = r.Scan(
			&locality,
			&amount,
			&municipality,
			&state,
			&lat,
			&long,
			&bnelat,
			&bnelong,
			&bswlat,
			&bswlong,
		)
		if err != nil {
			return nil, err
		}
		locs[idx] = locality
		states[idx] = state
		an[idx] = question.Answer{
			Text: fmt.Sprintf("%d", amount),
		}
		if bnelat == 0 {
			nelat[idx] = lat
			nelong[idx] = long
			swlat[idx] = lat
			swlong[idx] = long
		} else {
			nelat[idx] = bnelat
			nelong[idx] = bnelong
			swlat[idx] = bswlat
			swlong[idx] = bswlong
		}
		clat[idx] = lat
		clon[idx] = long
		idx++
	}

	correct := rand.Intn(4)

	cs := states[correct]
	if !strings.Contains(cs, "Territory") {
		cs = "state of " + cs
	}
	cl := locs[correct]
	if i := strings.IndexRune(cl, '('); i != -1 {
		cl = strings.Trim(cl[:i], " ")
	}

	qu := &question.Question{
		Answers: an,
		Correct: correct,
		Text:    fmt.Sprintf(q, cl, cs),
		Region: question.QuestionRegion{
			NELat:      nelat[correct],
			NELong:     nelong[correct],
			SWLat:      swlat[correct],
			SWLong:     swlong[correct],
			CenterLat:  clat[correct],
			CenterLong: clon[correct],
		},
	}

	return qu, nil

}
