package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

const queuePath = "../common/queue"
const scrapePath = "../common/scrape"

func parse() (string, error) {
	fScrape, err := os.Open(scrapePath)
	if err != nil {
		return "", err
	}
	defer fScrape.Close()

	var buffer bytes.Buffer
	buffer.WriteString("Good morning!!\n\n<ul>")

	scanner := bufio.NewScanner(fScrape)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "\t")
		buffer.WriteString("<li>" + s[1] + "\n" + s[0] + "</li>\n\n")
	}
	buffer.WriteString("</ul>")

	if err = scanner.Err(); err != nil {
		return "", err
	}

	if _, err := os.Stat(queuePath); err != nil {
		return buffer.String(), nil
	}

	fQueue, err := os.Open(queuePath)
	if err != nil {
		return "", err
	}

	buffer.WriteString("Bonus links!\n\n<ul>")
	scanner = bufio.NewScanner(fQueue)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ";")
		buffer.WriteString("<li>" + s[1] + "\n" + s[0] + "</li>\n\n")
	}
	buffer.WriteString("</ul>")

	if err = scanner.Err(); err != nil {
		return "", err
	}
	fQueue.Close()
	os.Remove(queuePath)
	return buffer.String(), nil
}
