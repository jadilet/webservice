package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jadilet/webservice/pkg/http/rest"
	"github.com/jadilet/webservice/pkg/service"
)

const OSRM_API_URL = "http://router.project-osrm.org/route/v1/driving/"

func main() {

	var (
		httpAddr = flag.String("port", "8080", "HTTP listen address")
	)
	flag.Parse()

	port := os.Getenv("PORT")

	if port == "" {
		port = *httpAddr
	}

	router := mux.NewRouter()
	apiService := service.NewOsrmService(OSRM_API_URL)
	routeService := service.NewRouteService(apiService)

	server := rest.NewRouteServer(routeService)

	router.HandleFunc("/routes", server.Route).Methods("GET")

	// Set up logging and panic recovery middleware.
	router.Use(func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(os.Stdout, h)
	})

	router.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))

	log.Printf("Starting HTTP server on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
