package main

import (
	"bufio"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"time"
)

func send(body string) {
	from, pass, to := initialize()
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Good morning! It's "+time.Now().Weekday().String())
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, from, pass)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	log.Print("sent")
}

func initialize() (string, string, string) {
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
	scanner.Scan()
	to := scanner.Text()
	return from, pass, to
}
