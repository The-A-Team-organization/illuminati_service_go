package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Compromised(t *testing.T) {

    server := httptest.NewServer(http.HandlerFunc(Compromised))
    defer server.Close()
    resp, _  := http.Get(server.URL)
	var payload struct {
	 	Password string `json:"password"`
	}
	json.NewDecoder(resp.Body).Decode(&payload)
    if payload.Password == "" {
        t.Errorf("Expected post@pon.com, post1@pon.com, post2@pon.com, got %s", payload.Password)
    }
}
	
