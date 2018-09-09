package question

type Answer struct {
	Text string `json:"text"`
}

type QuestionRegion struct {
	NELat      float64 `json:"nelat"`
	NELong     float64 `json:"nelon"`
	SWLat      float64 `json:"swlat"`
	SWLong     float64 `json:"swlon"`
	CenterLat  float64 `json:"lat"`
	CenterLong float64 `json:"lon"`
}

type Question struct {
	Text    string         `json:"text"`
	Answers [4]Answer      `json:"answers"`
	Correct int            `json:"correct"`
	Region  QuestionRegion `json:"region"`
}
