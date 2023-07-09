package main

import (
	"fmt"
	"github.com/go-mail/mail"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	m := mail.NewMessage()
	m.SetHeader("From", "team@paul-codes.com")
	m.SetHeader("To", os.Getenv("PERSONAL_EMAIL"))
	//m.SetAddressHeader("Cc", "oliver.doe@example.com", "Oliver")
	m.SetHeader("Subject", "Action")
	m.SetBody("text/html", "Howdy <b>Paul</b> testing the Go mail app, let me know how it goes... Get it!!!")
	//m.Attach("lolcat.jpeg")
	d := mail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("APP_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Printf("%s uses this password %s\n", os.Getenv("EMAIL"), os.Getenv("APP_PASSWORD"))
}
