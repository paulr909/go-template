package main

import (
	"fmt"
	"github.com/go-mail/mail"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type Form struct {
	Name, Email, Message string
}

func sendEmail(r *http.Request) {
	msg := Form{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	m := mail.NewMessage()
	m.SetHeader("From", "team@paul-codes.com")
	m.SetHeader("To", os.Getenv("PERSONAL_EMAIL"))
	m.SetHeader("Subject", "Contact Form")
	m.SetBody("text/html", fmt.Sprintf("Hi Paul,<br><br>New contact details below.<br><br>Name: %s <br>Email: %s <br>Message: %s", msg.Name, msg.Email, msg.Message))
	d := mail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("APP_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
