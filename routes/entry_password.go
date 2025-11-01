package routes

import (
	"encoding/json"
	"fmt"
	"illuminati/go/microservice/utils"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/XANi/loremipsum"
)


type participants struct {
	Participants []string `json:"participants"`
}


type entryPasswordService struct {
	emailSender utils.EmailSender
    participantsURL string
}

func NewEntryPasswordService(emailSender utils.EmailSender,  participantsURL string) *entryPasswordService {
	return &entryPasswordService{
		emailSender: emailSender,
		participantsURL : participantsURL,
	}
}

var (
	participants_url = os.Getenv("PARTICIPANTS_URL")
	es = NewEntryPasswordService(utils.SingletonEmailSender, participants_url)
)


func getRandomWord() string {

	loremIpsumGenerator := loremipsum.New()
	Word := loremIpsumGenerator.Word()
	log.Print("new word : ", Word)
	return strings.TrimSpace(Word)

}

func (es *entryPasswordService)getAppParticipants() ([]string, error) {

	resp, err := http.Get(es.participantsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	var data participants
	json.NewDecoder(resp.Body).Decode(&data)

	return data.Participants, nil
}

func (es *entryPasswordService)sendEntryPasswordEmail(word string, participants []string) error {

	text := fmt.Sprintf(`Hello,

		You're subscribed to our Latin vocabulary service, and today’s Word of the Day is %s.

		Learn its meaning, usage, and examples to expand your vocabulary!

		Happy learning
		— The Latin Words Team
		`,
		word)

	err := es.emailSender.SendEmail("Word of the Day", text, participants)
	if err != nil {
		log.Fatal("Something went wrong during sending the emails..")
		return err
	}

	return nil
}

func (es *entryPasswordService)getNewEntryPassword(w http.ResponseWriter, r *http.Request) {

	word := getRandomWord()
	log.Print("participants url : ", es.participantsURL)
	participants, err := es.getAppParticipants()
	if err != nil {
  		log.Fatal("Get no participants :", err)
	}
	log.Print("Got participants :", participants)

	if word == "" {
		log.Fatal("the word field is blank")
	}

	err = es.sendEntryPasswordEmail(word, participants)

	hashed, err := utils.HashPassword(word)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
  		return
	}

	data, err :=  utils.SerializePasswordHash(hashed)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
  		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}