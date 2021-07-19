package model

// OsrmResponse is a struct that contains the response from the OSRM API
type OsrmResponse struct {
	Code   string `json:"code"`
	Routes []struct {
		Distance float64 `json:"distance"`
		Duration float64 `json:"duration"`
		Legs     []struct {
			Distance    float64 `json:"distance"`
			Duration    float64 `json:"duration"`
			Destination string  `json:"destination"`
		} `json:"legs"`
	} `json:"routes"`
}

type Route struct {
	Distance    float64 `json:"distance"`
	Duration    float64 `json:"duration"`
	Destination string  `json:"destination"`
}

// RouteResponse is a struct that contains the response of web service
type RouteResponse struct {
	Source string `json:"source"`
	Routes []Route `json:"routes"`
}
