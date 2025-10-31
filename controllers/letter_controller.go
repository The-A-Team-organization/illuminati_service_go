package controllers

import (
	"illuminati/go/microservice/service"
	"log"
	"net/http"
)

func PostLetter(w http.ResponseWriter, r *http.Request){
	err := service.BuildLetterEmail(r)
	if err != nil {
		log.Fatal("Error :", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusAccepted)
}