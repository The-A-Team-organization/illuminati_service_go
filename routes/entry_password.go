package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"illuminati/go/microservice/utils"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/XANi/loremipsum"
)

var (
	username    = os.Getenv("EMAIL_USERNAME")
	password    = os.Getenv("EMAIL_PASSWORD")
	participantsURL = os.Getenv("PARTICIPANTS_URL")
)

func GetRandomWord() string {

	loremIpsumGenerator := loremipsum.New()
	Word := loremIpsumGenerator.Word()
	log.Print("new word : ", Word)
	return strings.TrimSpace(Word)

}

func GetAppParticipants(url string) ([]string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data struct {
		Participants []string `json:"participants"`
	}

	json.NewDecoder(resp.Body).Decode(&data)

	return data.Participants, nil
}

func SendEntryPasswordEmail(word string, participants []string) error {

	if word == "" {
		return errors.New("the word field is blank")
	}

	text := fmt.Sprintf(`Hello,

		You're subscribed to our Latin vocabulary service, and today’s Word of the Day is %s.

		Learn its meaning, usage, and examples to expand your vocabulary!

		Happy learning
		— The Latin Words Team
		`,
		word)

	err := utils.SingletonEmailSender.SendEmail("Word of the Day", text, participants)
	if err != nil {
		log.Fatal("Something went wrong during sending the emails..")
		return err
	}

	return nil
}

func GetNewEntryPassword(w http.ResponseWriter, r *http.Request) {
	
	
	newWord := GetRandomWord()
	log.Print("participants url : ", participantsURL)
	participants, err := GetAppParticipants(participantsURL)
	if err != nil {
		log.Print("Get no participants :", err)
		w.WriteHeader(http.StatusInternalServerError)
  		return
	}
	log.Print("Got participants :", participants)
	
	SendEntryPasswordEmail(newWord,participants)

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