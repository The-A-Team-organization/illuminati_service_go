package routes

import (
	"encoding/json"
	"illuminati/go/microservice/utils"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
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
	ls = NewLetterService(utils.GetInstance())
)

type Letter struct {
	Topic        string   `json:"topic"`
	Text         string   `json:"text"`
	TargetEmails []string `json:"target_emails"`
}

func (ls *LetterService) SendLetterEmail(request *http.Request) error {
	var g errgroup.Group
	var letter Letter
	err := json.NewDecoder(request.Body).Decode(&letter)
	if err != nil {
		log.Println("Something went wrong while parsing the request body...")
		return err
	}

    g.Go(
		func() error {
		err = ls.emailSender.SendEmail(letter.Topic, letter.Text, letter.TargetEmails)
		return err
		},
	)
	if err := g.Wait(); err != nil {
		log.Println("Something went wrong while sending the emails...")
		return err
	}

	return nil
}

func (ls *LetterService)PostLetter(w http.ResponseWriter, r *http.Request) {

	err := ls.SendLetterEmail(r)
	if err != nil {
		log.Print("Error :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}


