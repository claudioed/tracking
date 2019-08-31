package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
	fmt.Fprintf(w, "UP %s %s ", serviceName, serviceVersion)
}

func reqOtherService() string {
	resp, err := http.Get(os.Getenv("TARGET_SERVICE_HOST"))
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	return string(body)
}
