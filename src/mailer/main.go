package main

import (
	"fmt"
	"os/exec"
	"time"
)

func every(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

func parseAndSend() {
	_, err := exec.Command("python3", "../scraper/scraper.py").Output()
	if err != nil {
		fmt.Println("%s", err)
	}

	str, err := parse()
	if err != nil {
		panic(err)
	}
	send(str)
}

func main() {
	every(5*time.Second, parseAndSend)
}
