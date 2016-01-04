package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

const returnChar = 13
const spaceChar = 32
const rChar = 114

func replaceText(msgs []string, extras ...string) {
	clearDisplay()
	for _, s := range msgs {
		if s == "" {
			continue
		}
		println(s)
	}
	for _, s := range extras {
		if s == "" {
			continue
		}
		println(s)
	}
}

func clearDisplay() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func inSeconds(d time.Duration) string {
	d = ((d / time.Second) * time.Second)
	if d == 0 {
		return "0s"
	}
	return ((d / time.Second) * time.Second).String()
}

func playSound() {
	cmd := exec.Command("paplay", "clap.wav")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
