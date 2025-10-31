package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_Compromised(t *testing.T) {
	test_backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer test_backend.Close()
	os.Setenv("PARTICIPANTS_URL", test_backend.URL)
    server := httptest.NewServer(http.HandlerFunc(WordSender))
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

func Test_WhenPartisipantUrlIsNULL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(WordSender))
    defer server.Close()
    resp, _  := http.Get(server.URL)
	if resp.StatusCode == http.StatusOK {
        t.Errorf("Expected error here")
    }
	
}
