package main

import (
	"illuminati/go/microservice/routes"
	"log"
	"net/http"
)

func main() {

	routes.SetupAPI()
	log.Fatal(http.ListenAndServe(":1080", nil))
}
