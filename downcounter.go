package main

import "time"

type downcounter struct {
	Title   string
	Minutes int
	Seconds int
	counterData
}

func (c *downcounter) count(quiet bool) {
	c.start = time.Now()
	for f := countdown; f != nil; {
		f = f(c)
		time.Sleep(time.Second)
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

type downcounterFSM func(*downcounter) downcounterFSM

func countdown(c *downcounter) downcounterFSM {
	if c.elapsed() > c.total() {
		return nil
	}
	replaceText(c.Title, inSeconds(c.elapsed()), inSeconds(c.remaining()))
	select {
	case b := <-stdinChars:
		switch b {
		case spaceChar:
			c.pauseStart = time.Now()
			return pausedown
		case rChar:
			c.start = time.Now()
			return countdown
		case returnChar:
			return nil
		default:
			return countdown
		}
	default:
		return countdown
	}
}

func pausedown(c *downcounter) downcounterFSM {
	replaceText(c.Title, inSeconds(c.elapsed()), inSeconds(c.remaining()), "PAUSED")
	select {
	case b := <-stdinChars:
		switch b {
		case spaceChar:
			c.pauses = append(c.pauses, c.pauseElapsed())
			return countdown
		case rChar:
			c.start = time.Now()
			return pausedown
		case returnChar:
			return nil
		default:
			return countdown
		}
	default:
		return pausedown
	}
}
