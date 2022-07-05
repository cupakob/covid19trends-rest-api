package domain

type Response struct {
	Global    Global
	Countries []Country
	Date      string
}

type Global struct {
	NewConfirmed   int
	TotalConfirmed int
	NewDeaths      int
	TotalDeaths    int
	NewRecovered   int
	TotalRecovered int
}

type Country struct {
	Country        string
	CountryCode    string
	Slug           string
	NewConfirmed   int
	TotalConfirmed int
	NewDeaths      int
	TotalDeaths    int
	NewRecovered   int
	TotalRecovered int
	Date           string
}
