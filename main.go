package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var port = 8080
	serviceName := os.Getenv("SERVICE_NAME")
	fmt.Printf("Running on :%d", port)
	path := fmt.Sprintf("/%s", serviceName)
	healthPath := fmt.Sprintf("/%s/health", serviceName)
	http.HandleFunc(path, handler)
	http.HandleFunc(healthPath, health)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	serviceName := os.Getenv("SERVICE_NAME")
	serviceVersion := os.Getenv("SERVICE_VERSION")
	fmt.Fprintf(w, "service name %s version %s", serviceName, serviceVersion)
}

func health(w http.ResponseWriter, r *http.Request) {
	serviceName := os.Getenv("SERVICE_NAME")
	serviceVersion := os.Getenv("SERVICE_VERSION")
	fmt.Fprintf(w, "UP %s %s", serviceName, serviceVersion)
}
