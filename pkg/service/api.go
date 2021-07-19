package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/jadilet/webservice/pkg/model"
)

type osrmService struct {
	url string
}

// OSRM API service
type ApiService interface {
	Route(src string, dst []string) (model.OsrmResponse, error)
}

func NewOsrmService(url string) ApiService {
	return &osrmService{url}
}

// Route takes a source and a destination and returns a route between them
func (s *osrmService) Route(src string, dst []string) (model.OsrmResponse, error) {
	url := fmt.Sprintf("%s%s", s.url, src)
	for _, d := range dst {
		url = fmt.Sprintf("%s;%s", url, d)
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return model.OsrmResponse{}, err
	}

	q := req.URL.Query()

	q.Add("overview", "false")

	req.URL.RawQuery = q.Encode()

	log.Println("Sending request to OSRM API:", url)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return model.OsrmResponse{}, err
	}

	defer resp.Body.Close()

	var route model.OsrmResponse
	err = json.NewDecoder(resp.Body).Decode(&route)

	if err != nil {
		return model.OsrmResponse{}, err
	}

	if route.Code != "Ok" {
		return model.OsrmResponse{}, fmt.Errorf("OSRM returned an error: %s", route.Code)
	}

	sort.Slice(route.Routes, func(i, j int) bool {
		return route.Routes[i].Duration < route.Routes[j].Duration
	})

	sort.Slice(route.Routes, func(i, j int) bool {
		return route.Routes[i].Duration == route.Routes[j].Duration &&
			route.Routes[i].Distance < route.Routes[j].Distance
	})

	return route, nil
}
