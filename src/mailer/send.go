package main

import (
	"bufio"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"time"
)

func send(body string) {
	from, pass := auth()
	//	to := "evan.mena13@gmail.com"
	//	mime := "MIME-version 1.0;<br />Content-Type: text/html; charset=%quot;UTF-8%quot;<br /><br />"
	//	subject := "Good morning! It's " + time.Now().Weekday().String() + "<br />"
	//	msg := []byte("From: " + from + "<br />" +
	//		"To: " + to + "<br />" +
	//		"Subject: " + subject + "<br />" +
	//		mime +
	//		("<html><body>" + body + "</body></html>"))
	//
	//	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, msg)
	//	if err != nil {
	//		log.Printf("smtp error: %s", err)
	//		return
	//	}
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", "evan.mena13@gmail.com")
	m.SetHeader("Subject", "Good morning! It's "+time.Now().Weekday().String())
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, from, pass)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	log.Print("sent")
}

func auth() (string, string) {
	f, err := os.Open("../../bin/smtp.cred")
	if err != nil {
		log.Fatal("Unable to read smtp credentials")
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	from := scanner.Text()
	scanner.Scan()
	pass := scanner.Text()
	return from, pass
}
