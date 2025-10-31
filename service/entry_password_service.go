package service

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
	 username = os.Getenv("EMAIL_USERNAME") 
	 password = os.Getenv("EMAIL_PASSWORD")
)

func GetRandomWord() string {

	loremIpsumGenerator := loremipsum.New()
	Word := loremIpsumGenerator.Word()
	log.Print("new word : ", Word)
	return strings.TrimSpace(Word)

}

func GetAppParticipants(url string) ([]string, error) {

	resp, err  := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	
	var data struct {
	 	Participants []string `json:"participants"`
	}

	json.NewDecoder(resp.Body).Decode(&data);

	return data.Participants, nil
}

func BuildEntryPasswordEmail(word string, participants []string) error{

	if(word == "") {
		return errors.New("the word field is blank")
	}

	text := fmt.Sprintf(`Hello,

		You're subscribed to our Latin vocabulary service, and today’s Word of the Day is %s.

		Learn its meaning, usage, and examples to expand your vocabulary!

		Happy learning
		— The Latin Words Team
		`, 
	word)

	err := utils.SendEmail("Word of the Day", text, participants)
	if err != nil {
		log.Fatal("Something went wrong during sending the emails..")
		return err
	}

	return nil
}
