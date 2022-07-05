package domain

type Output struct {
	Cases     int    `json:"cases"`
	CasesNew  int    `json:"cases_new"`
	Deaths    int    `json:"deaths"`
	DeathsNew int    `json:"deaths_new"`
	Timestamp int64  `json:"timestamp"`
	Date      string `json:"date"`
	DaysPast  int    `json:"days_past"`
}
