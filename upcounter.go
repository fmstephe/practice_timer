package main

import "time"

type upcounter struct {
	Title string
	// Separate these mutable fields
	start      time.Time
	pauseStart time.Time
	pauses     []time.Duration
}

func (c *upcounter) count() {
	c.start = time.Now()
	for f := countup; f != nil; {
		f = f(c)
		time.Sleep(time.Second)
	}
}

func (c *upcounter) elapsed() time.Duration {
	return time.Now().Sub(c.start)
}

func (c *upcounter) totalPaused() time.Duration {
	var paused time.Duration
	for _, p := range c.pauses {
		paused += p
	}
	return paused
}

func (c *upcounter) pauseElapsed() time.Duration {
	return time.Now().Sub(c.pauseStart)
}

func (c *upcounter) restart() {
	c.start = time.Now()
	c.pauseStart = c.start
	c.pauses = nil
}

type upcounterFSM func(*upcounter) upcounterFSM

func countup(c *upcounter) upcounterFSM {
	replaceText(c.Title, inSeconds(c.elapsed()))
	select {
	case b := <-stdinChars:
		switch b {
		case spaceChar:
			c.pauseStart = time.Now()
			return pauseup
		case rChar:
			c.restart()
			return countup
		case returnChar:
			return nil
		default:
			return countup
		}
	default:
		return countup
	}
}

func pauseup(c *upcounter) upcounterFSM {
	replaceText(c.Title, inSeconds(c.elapsed()), "PAUSED")
	select {
	case b := <-stdinChars:
		switch b {
		case spaceChar:
			c.pauses = append(c.pauses, c.pauseElapsed())
			return countup
		case rChar:
			c.restart()
			return pauseup
		case returnChar:
			return nil
		default:
			return countup
		}
	default:
		return pauseup
	}
}
