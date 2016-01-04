package main

import "time"

type counter interface {
	display() []string
	restart()
	addPause(time.Duration)
	finished() bool
	finish()
}

func runFSM(c counter) {
	c.restart()
	tick := time.Now()
	for f := countFSM; f != nil; {
		gap := time.Now().Sub(tick)
		tick = time.Now()
		f = f(c, gap)
		time.Sleep(100 * time.Millisecond)
	}
	c.finish()
}

type counterFSM func(counter, time.Duration) counterFSM

func countFSM(c counter, gap time.Duration) counterFSM {
	if c.finished() {
		return nil
	}
	replaceText(c.display())
	select {
	case b := <-stdinChars:
		switch b {
		case spaceChar:
			return pauseFSM
		case rChar:
			c.restart()
			return countFSM
		case returnChar:
			return nil
		default:
			return countFSM
		}
	default:
		return countFSM
	}
}

func pauseFSM(c counter, gap time.Duration) counterFSM {
	c.addPause(gap)
	replaceText(c.display(), "PAUSED")
	select {
	case b := <-stdinChars:
		switch b {
		case spaceChar:
			return countFSM
		case rChar:
			c.restart()
			return pauseFSM
		case returnChar:
			return nil
		default:
			return countFSM
		}
	default:
		return pauseFSM
	}
}
