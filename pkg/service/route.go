package service

import (
	"github.com/jadilet/webservice/pkg/model"
)

type routeService struct {
	api ApiService
}

type RouteService interface {
	GetRoute(src string, dst []string) (*model.RouteResponse, error)
}

func NewRouteService(api ApiService) RouteService {
	return &routeService{api}
}

func (s *routeService) GetRoute(src string, dst []string) (*model.RouteResponse, error) {
	resp, err := s.api.Route(src, dst)

	if err != nil {
		return nil, err
	}

	var r model.RouteResponse
	r.Source = src

	for _, route := range resp.Routes {
		for i, leg := range route.Legs {
			r.Routes = append(r.Routes, model.Route{Distance: leg.Distance,
				Duration:    leg.Duration,
				Destination: dst[i]})
		}
	}

	return &r, nil
}
