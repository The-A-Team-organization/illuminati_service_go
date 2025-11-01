package main

type Letter struct {
	Topic        string   `json:"topic"`
	Text         string   `json:"text"`
	TargetEmails []string `json:"target_emails"`
}