package service

import (
	"encoding/json"
	"illuminati/go/microservice/utils"
	"log"
	"net/http"
)

func BuildLetterEmail(request *http.Request) error{

	var letter struct {
		Topic        string `json:"topic"`
		Text         string `json:"text"`
		TargetEmails []string `json:"target_emails"`
	}

	err := json.NewDecoder(request.Body).Decode(&letter)
	if err != nil {
		log.Fatal("Something went wrong while parsing the request body...")
		return err
	}

	err = utils.SendEmail(letter.Topic, letter.Text, letter.TargetEmails)
	if err != nil {
		log.Fatal("Something went wrong while sending the emails...")
		return err
	}

	return nil
}