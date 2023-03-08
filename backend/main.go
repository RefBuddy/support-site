package main

import (
	"fmt"
	"net/smtp"
)

const emailSender = "duckwxrth@gmail.com"
const emailPassword = "zipgmipuboimouse"
const emailHost = "smtp.gmail.com"
const emailPort = "587"

type ContactForm struct {
	Name    string
	Email   string
	Message string
}

func main() {
	form := &ContactForm{
		Name:    "John Doe",
		Email:   "johndoe@example.com",
		Message: "Hello, I am sending this message from Ref Buddy contact form.",
	}

	to := []string{"duckwxrth@gmail.com"}
	message := []byte("To: " + to[0] + "\r\n" +
		"Subject: Message from Ref Buddy Contact Form\r\n" +
		"\r\n" +
		form.Message + "\r\n")

	auth := smtp.PlainAuth("", emailSender, emailPassword, emailHost)
	err := smtp.SendMail(emailHost+":"+emailPort, auth, emailSender, to, message)
	if err != nil {
		fmt.Printf("Failed to send email: %s", err)
	}
}
