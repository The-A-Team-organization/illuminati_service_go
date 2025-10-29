package routes

import (
	"illuminati/go/microservice/controllers"
	"net/http"
)

func SetupAPI(){
	http.HandleFunc("/compromised", controllers.Compromised)
}

