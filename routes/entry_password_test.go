package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func Test_GetRandomWord(t *testing.T) {
    word := GetRandomWord()
    if strings.TrimSpace(word) == "" {
        t.Error("Expected non-empty word")
    }
}

func Test_GetAppParticipants(t *testing.T) {

    test :=  func(w http.ResponseWriter, r *http.Request) {
        type data struct {
	 	Participants []string `json:"participants"`}
        var example data 
        example.Participants = []string{"test1@gmail.com", "test123@gmail.com"}
        dataSend, _  := json.Marshal(example)
        fmt.Println(dataSend)
        w.Header().Set("Content-Type", "application/json")
        w.Write(dataSend)
    }  
    server := httptest.NewServer(http.HandlerFunc(test))
    defer server.Close()
    resp, _  := GetAppParticipants(server.URL)
    if resp == nil {
        t.Errorf("Expected post@pon.com, post1@pon.com, post2@pon.com, got %s", resp)
    }
}

func Test_SendWordEmail(t *testing.T) {

    err := SendEntryPasswordEmail("", []string{})
    if err == nil{
         t.Errorf("Here should be err")
    } 
    err = SendEntryPasswordEmail("Lorem", []string{})
    if err != nil{
         t.Errorf("Here shouldn`t be err")
    } 
    
}

func Test_GetNewEntryPassword(t *testing.T) {
	test_backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer test_backend.Close()
	os.Setenv("PARTICIPANTS_URL", test_backend.URL)
    server := httptest.NewServer(http.HandlerFunc(GetNewEntryPassword))
    defer server.Close()
    resp, _  := http.Get(server.URL)
	var payload struct {
	 	Password string `json:"entry_password"`
	}
	json.NewDecoder(resp.Body).Decode(&payload)
    if payload.Password == "" {
        t.Errorf("Expected some hashed password, got %s", payload.Password)
    }
}

func Test_WhenPartisipanceUrlIsNULL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(GetNewEntryPassword))
    defer server.Close()
    resp, _  := http.Get(server.URL)
	if resp.StatusCode == http.StatusOK {
        t.Errorf("Expected error here")
    }
	
}
