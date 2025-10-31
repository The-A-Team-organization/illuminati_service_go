package routes

import (
	"illuminati/go/microservice/controllers"
	"net/http"
)

func SetupAPI(){
	http.HandleFunc("/entry_password", controllers.GetNewEntryPassword)
	http.HandleFunc("/send_letter", controllers.PostLetter)
}

