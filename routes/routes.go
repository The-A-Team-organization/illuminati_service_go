package routes

import (
	"illuminati/go/microservice/controllers"
	"net/http"
)

func SetupAPI(){
	http.HandleFunc("/new-word", controllers.WordResender)
}

