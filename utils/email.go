package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/wneessen/go-mail"
)

func SendEmail(topic, text string, targetEmails []string) error{

	username := os.Getenv("EMAIL_USERNAME") 
	password := os.Getenv("EMAIL_PASSWORD")

	
	client, err := mail.NewClient(
		"smtp.gmail.com",
		mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover),
		mail.WithUsername(username),
		mail.WithPassword(password),
	)
	if err != nil {
		fmt.Println("Failed to create mail client:", err)
		return err
	}
	msg := mail.NewMsg()

	msg.Subject(topic)

	msg.SetBodyString(mail.TypeTextPlain, text)

	for _, email := range targetEmails {
		msg.From(username)
		msg.To(email)

		if err := client.DialAndSend(msg); err != nil {
			fmt.Println("Failed to send mail:", err)
			return err
		}

		log.Println("Email sent successfully!")
	}

	return nil
}
