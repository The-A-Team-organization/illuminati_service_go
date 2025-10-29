package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Compromised(t *testing.T) {

    server := httptest.NewServer(http.HandlerFunc(WordSender))
    defer server.Close()
    resp, _  := http.Get(server.URL)
	var payload struct {
	 	Password string `json:"password"`
	}
	json.NewDecoder(resp.Body).Decode(&payload)
    if payload.Password == "" {
        t.Errorf("Expected some hashed password, got %s", payload.Password)
    }
}
	
