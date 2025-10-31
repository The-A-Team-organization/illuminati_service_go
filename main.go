package main

import (
	"illuminati/go/microservice/routes"
	"log"
	"net/http"
)

func main() {
	routes.SetupAPI()
	log.Print("Starting a server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
