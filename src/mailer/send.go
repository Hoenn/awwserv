package main

import (
	"bufio"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func send(body string) {
	from, pass := auth()
	to := "evan.mena13@gmail.com"

	msg := "From: Awwserv\n" +
		"To: " + to + "\n" +
		"Subject: Good morning!\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("smtp error: %s", err)
		return
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
