package main

import (
	"github.com/munusamy/redditnews"
	"log"
	"net/smtp"
)

func main() {
	to := "munuswamymp@gmail.com"
	subject := "Go articles on Reddit"
	message := redditnews.Email()

	body := "To: " + to + "\r\nSubject: " +
		subject + "\r\n\r\n" + message

	auth := smtp.PlainAuth("", "satish.talim", "password", "smtp.gmail.com")
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"satish.talim@gmail.com",
		[]string{to},
		[]byte(body))
	if err != nil {
		log.Fatal("SendMail: ", err)
		return
	}
}
