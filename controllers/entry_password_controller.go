package controllers

import (
	"illuminati/go/microservice/service"
	"illuminati/go/microservice/utils"
	"log"
	"net/http"
	"os"
)


func GetNewEntryPassword(w http.ResponseWriter, r *http.Request) {
	participantsURL := os.Getenv("PARTICIPANTS_URL")
	
	newWord := service.GetRandomWord()
	log.Print("participants url : ", participantsURL)
	participants, err := service.GetAppParticipants(participantsURL)
	if err != nil {
		log.Print("Get no participants :", err)
		w.WriteHeader(http.StatusInternalServerError)
  		return
	}
	log.Print("Got participants :", participants)
	
	service.BuildEntryPasswordEmail(newWord,participants)

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