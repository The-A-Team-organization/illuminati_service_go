package routes

import (
	"net/http"
)

func SetupAPI(){
	http.HandleFunc("/entry_password", GetNewEntryPassword)
	http.HandleFunc("/send_letter", PostLetter)
}

