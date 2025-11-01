package routes

import (
	"net/http"
)

func SetupAPI(){
	http.HandleFunc("/entry_password", es.getNewEntryPassword)
	http.HandleFunc("/send_letter", ls.PostLetter)
}

