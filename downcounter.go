package main

import (
	"log"
	"os/exec"
	"time"
)

const spaceChar = 32
const rChar = 114

type downcounter struct {
	Title   string
	Minutes int
	Seconds int
	// Separate these mutable fields
	start  time.Time
	paused time.Duration
}

func (c *downcounter) count(quiet bool) {
	c.start = time.Now()
	for !c.isFinished() {
		c.updateDisplay()
		c.checkInput()
		time.Sleep(time.Second)
	}
	c.completeDisplay()
	if !quiet {
		c.playSound()
	}
}

func (c *downcounter) isFinished() bool {
	return c.elapsed() >= c.total()
}

func (c *downcounter) updateDisplay() {
	replaceText(c.Title, inSeconds(c.elapsed()), inSeconds(c.remaining()))
}

func (c *downcounter) completeDisplay() {
	replaceText(c.Title, inSeconds(c.total()), inSeconds(c.total()))
}

func (c *downcounter) total() time.Duration {
	m := time.Duration(c.Minutes) * time.Minute
	s := time.Duration(c.Seconds) * time.Second
	return m + s
}

func (c *downcounter) elapsed() time.Duration {
	return time.Now().Sub(c.start)
}

func (c *downcounter) remaining() time.Duration {
	return c.total() - c.elapsed() + time.Second
}

func (c *downcounter) checkInput() {
	select {
	case c := <-stdinChars:
		switch {
		case c == spaceChar:
			println("Space bar")
		case c == rChar:
			println("r")
		}
	default:
	}
}

func (c *downcounter) playSound() {
	clearDisplay()
	cmd := exec.Command("paplay", "clap.wav")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
