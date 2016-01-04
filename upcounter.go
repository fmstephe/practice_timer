package main

import (
	"log"
	"os/exec"
	"time"
)

type upcounter struct {
	Title   string
	Minutes int
	Seconds int
	// Separate these mutable fields
	start  time.Time
	paused time.Duration
}

func (c *upcounter) count(quiet bool) {
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

func (c *upcounter) isFinished() bool {
	return c.elapsed() >= c.total()
}

func (c *upcounter) updateDisplay() {
	clearDisplay()
	if c.Title != "" {
		println(c.Title)
	}
	println(inSeconds(c.elapsed()))
	println(inSeconds(c.remaining()))
}

func (c *upcounter) completeDisplay() {
	clearDisplay()
	if c.Title != "" {
		println(c.Title)
	}
	println(inSeconds(c.total()))
	println(inSeconds(c.total()))
}

func (c *upcounter) total() time.Duration {
	m := time.Duration(c.Minutes) * time.Minute
	s := time.Duration(c.Seconds) * time.Second
	return m + s
}

func (c *upcounter) elapsed() time.Duration {
	return time.Now().Sub(c.start)
}

func (c *upcounter) remaining() time.Duration {
	return c.total() - c.elapsed() + time.Second
}

func (c *upcounter) checkInput() {
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

func (c *upcounter) playSound() {
	clearDisplay()
	cmd := exec.Command("paplay", "clap.wav")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
