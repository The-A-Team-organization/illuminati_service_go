package utils

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/wneessen/go-mail"
)

type EmailSender interface {
	SendEmail(topic, text string, targetEmails []string) error
}

type emailSender struct {
	username string
	client *mail.Client
}

var (
	email_username       = os.Getenv("EMAIL_USERNAME")
	email_password       = os.Getenv("EMAIL_PASSWORD")
	email_host           = "smtp.gmail.com"
	singletonEmailSender *emailSender
	lock = &sync.Mutex{}
)

func init() {
	client, err := mail.NewClient(
		email_host,
		mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover),
		mail.WithUsername(email_username),
		mail.WithPassword(email_password),
	)
	if err != nil{
		log.Fatal("Cant create new SMTP clent :", err)
	}
	singletonEmailSender = &emailSender{
		email_username, 
		client,
	}
	log.Print("Initializated SMTP client!!!")
}

func GetInstance() *emailSender {
	return singletonEmailSender
}


func (es *emailSender) SendEmail(topic, text string, targetEmails []string) error {

	lock.Lock()
	defer lock.Unlock()
	msg := mail.NewMsg()

	msg.Subject(topic)

	msg.SetBodyString(mail.TypeTextPlain, text)

	for _, email := range targetEmails {
		msg.From(es.username)
		msg.To(email)

		if err := es.client.DialAndSend(msg); err != nil {
			fmt.Print("Failed to send mail:", err)
			return err
		}

	}
	log.Print("Emails were sent successfully!")

	return nil
}


