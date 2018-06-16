package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

func parse() (string, error) {
	fScrape, err := os.Open("../common/scrape")
	if err != nil {
		return "", err
	}
	defer fScrape.Close()

	var buffer bytes.Buffer
	buffer.WriteString("Good morning!!\n\n")

	scanner := bufio.NewScanner(fScrape)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "\t")
		buffer.WriteString(s[1] + "\n" + s[0] + "\n\n")
	}

	if err = scanner.Err(); err != nil {
		return "", err
	}

	fQueue, err := os.Open("../common/queue")
	if err != nil {
		return "", err
	}
	defer fQueue.Close()

	buffer.WriteString("Bonus links!\n\n")
	buffer.WriteString("~*~*~*~*~*~*~*~*~*~*~*~*~")

	scanner = bufio.NewScanner(fQueue)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ";")
		buffer.WriteString(s[1] + "\n" + s[0] + "\n\n")
	}

	if err = scanner.Err(); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
