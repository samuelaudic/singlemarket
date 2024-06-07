package utils

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to, subject, body string) error {
	from := "sender@example.com"
	password := ""

	// Setup server
	smtpHost := "localhost"
	smtpPort := "1025"

	// Setup message
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Send email
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
