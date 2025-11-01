package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/wneessen/go-mail"
)

type EmailSender interface {
	SendEmail(topic, text string, targetEmails []string) error
}

type emailSender struct {
	username string
	client *mail.Client
}

func NewEmailSender(username string, client *mail.Client) EmailSender {	
	return &emailSender{
		username: username,
		client: client,
	}
}


var (
	email_username    = os.Getenv("EMAIL_USERNAME")
	email_password    = os.Getenv("EMAIL_PASSWORD")
	email_host        = "smtp.gmail.com"
	client, err = mail.NewClient(
		email_host,
		mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover),
		mail.WithUsername(email_username),
		mail.WithPassword(email_password),
	)
	SingletonEmailSender = NewEmailSender(email_username, client) 
)



func (es *emailSender) SendEmail(topic, text string, targetEmails []string) error {

	msg := mail.NewMsg()

	msg.Subject(topic)

	msg.SetBodyString(mail.TypeTextPlain, text)

	for _, email := range targetEmails {
		msg.From(es.username)
		msg.To(email)

		if err := es.client.DialAndSend(msg); err != nil {
			fmt.Println("Failed to send mail:", err)
			return err
		}

		log.Println("Email sent successfully!")
	}

	return nil
}


