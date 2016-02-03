package fsm

import (
	"os"
	"os/exec"
)

const aChar = 97
const dChar = 100
const cChar = 99
const qChar = 113
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
