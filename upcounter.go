package main

import "time"

type upcounter struct {
	Title string
	// Separate these mutable fields
	start      time.Time
	paused     time.Duration
	isFinished bool
}

func (c *upcounter) count(quiet bool) {
	c.start = time.Now()
	for !c.isFinished {
		c.checkInput()
		replaceText(c.Title, inSeconds(c.elapsed()))
		time.Sleep(time.Second)
	}
	replaceText(c.Title, inSeconds(c.elapsed()))
}

func (c *upcounter) elapsed() time.Duration {
	return time.Now().Sub(c.start) - c.paused
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
