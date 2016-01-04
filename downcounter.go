package main

import "time"

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
	for c.elapsed() < c.total() {
		replaceText(c.Title, inSeconds(c.elapsed()), inSeconds(c.remaining()))
		c.checkInput()
		time.Sleep(time.Second)
	}
	replaceText(c.Title, inSeconds(c.total()), inSeconds(c.total()))
	if !quiet {
		clearDisplay()
		playSound()
	}
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
