package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func replaceText(strings ...string) {
	clearDisplay()
	for _, s := range strings {
		println(s)
	}
}

func clearDisplay() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func inSeconds(d time.Duration) string {
	if d == 0 {
		return "0s"
	}
	return ((d / time.Second) * time.Second).String()
}

func playSound() {
	clearDisplay()
	cmd := exec.Command("paplay", "clap.wav")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
