package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/jadilet/webservice/pkg/service"
)

type routeServer struct {
	routeService service.RouteService
	regex        *regexp.Regexp // validation regexp for latitude and longitude
}

func NewRouteServer(routeService service.RouteService) *routeServer {
	return &routeServer{
		routeService: routeService,
		regex:        regexp.MustCompile(`^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`),
	}
}

// renderJSON renders 'v' as JSON and writes it as a response into w.
func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *routeServer) Route(w http.ResponseWriter, r *http.Request) {

	src := r.URL.Query().Get("src")
	dst := r.URL.Query()["dst"]

	// validate src and dst
	if len(src) == 0 || len(dst) == 0 {
		http.Error(w, "src and dst are mandatory", http.StatusBadRequest)
		return
	}

	// Validate src and dst coords
	if !s.regex.MatchString(src) {
		http.Error(w, "Source invalid coordinates", http.StatusBadRequest)
		return
	}

	for _, v := range dst {
		if !s.regex.MatchString(v) {
			http.Error(w, "Destination invalid coordinates", http.StatusBadRequest)
			return
		}
	}

	log.Printf("Routing from %s to %v", src, dst)

	routes, err := s.routeService.GetRoute(src, dst)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(w, routes)
}
