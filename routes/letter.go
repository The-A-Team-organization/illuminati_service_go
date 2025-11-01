package routes

import (
	"encoding/json"
	"illuminati/go/microservice/utils"
	"log"
	"net/http"
)

type LetterService struct {
	emailSender utils.EmailSender
}

func NewLetterService(emailSender utils.EmailSender) *LetterService {
	return &LetterService{
		emailSender: emailSender,
	}
}

var(
	ls = NewLetterService(utils.SingletonEmailSender)
)

type Letter struct {
	Topic        string   `json:"topic"`
	Text         string   `json:"text"`
	TargetEmails []string `json:"target_emails"`
}

func (ls *LetterService) SendLetterEmail(request *http.Request) error {
	var letter Letter
	err := json.NewDecoder(request.Body).Decode(&letter)
	if err != nil {
		log.Fatal("Something went wrong while parsing the request body...")
		return err
	}

	err = ls.emailSender.SendEmail(letter.Topic, letter.Text, letter.TargetEmails)
	if err != nil {
		log.Fatal("Something went wrong while sending the emails...")
		return err
	}

	return nil
}

func PostLetter(w http.ResponseWriter, r *http.Request) {

	err := ls.SendLetterEmail(r)
	if err != nil {
		log.Fatal("Error :", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusAccepted)
}


