package controllers

import (
	"illuminati/go/microservice/service"
	"illuminati/go/microservice/utils"
	"net/http"
	"os"
)


var mockURL  = os.Getenv("MOCK_URL")

func Compromised(w http.ResponseWriter, r *http.Request) {
	
	newWord := service.GetRandomWord()
	participants, _ := service.GetAppParticipants(mockURL)
	service.SendWordEmail(newWord,participants)

	hashed, err := utils.HashPassword(newWord)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
  		return
	}

	data, err :=  utils.SerializePasswordHash(hashed)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
  		return
	}

	w.Write(data)
}