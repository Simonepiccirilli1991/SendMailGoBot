package service

import (
	"SendMailGoBot/model"
	"fmt"

	"gopkg.in/gomail.v2"
)

func sendMail(dto model.MailDto) {
	// Create a new message
	m := gomail.NewMessage()

	// Set the from, to, subject, and body
	m.SetHeader("From", dto.From)
	m.SetHeader("To", "hardcodedmail@mail.com")
	m.SetHeader("Subject", "Hello from Go Bot!")
	m.SetBody("text/plain", dto.Testo)

	// Create a new SMTP dialer to connect to an SMTP server
	d := gomail.NewDialer("smtp.example.com", 587, "sender@example.com", "password")

	// Send the email using the SMTP dialer
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Failed to send email:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
