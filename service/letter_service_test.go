package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func Test_BuildLetterEmail(t *testing.T) {
    type letter struct {
	Topic        string `json:"topic"`
	Text         string `json:"text"`
	TargetEmails []string `json:"target_emails"`}
	
    var example letter 
    example.Topic = "New Letter"
	example.Text = "Here is new letter"
	example.TargetEmails = []string{"test1@gmail.com", "test123@gmail.com"}

    dataSend, _  := json.Marshal(example)

    fmt.Println(dataSend)
  	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(dataSend))

	err := BuildLetterEmail(request)
    if err != nil {
        t.Error("Expected post@pon.com, post1@pon.com, post2@pon.com, got")
    }
	
}