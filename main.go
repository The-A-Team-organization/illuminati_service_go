package main

import (
	"illuminati/go/microservice/controllers"
	"illuminati/go/microservice/routes"
	"log"
	"net/http"
)

func main() {

	var mutexManager controllers.MutexManager
	routes.SetupAPI(mutexManager)
	log.Fatal(http.ListenAndServe(":1080", nil))
}
