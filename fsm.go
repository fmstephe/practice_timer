package main

import "time"

type counter interface {
	display() []string
	restart()
	addElapsed(time.Duration)
	addPause(time.Duration)
	finish() *CounterRecord
	cancel()
	finished() bool
}

func runFSM(c counter) *CounterRecord {
	c.restart()
	tick := time.Now()
	f := countFSM
	for !c.finished() {
		gap := time.Now().Sub(tick)
		tick = time.Now()
		f = f(c, gap)
		time.Sleep(100 * time.Millisecond)
	}
	return c.finish()
}

type counterFSM func(counter, time.Duration) counterFSM

func countFSM(c counter, gap time.Duration) counterFSM {
	c.addElapsed(gap)
	replaceText(c.display())
	select {
	case b := <-stdinChars:
		switch b {
		case spaceChar:
			return pauseFSM
		case rChar:
			c.restart()
			return countFSM
		case cChar:
			c.cancel()
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
		case cChar:
			c.cancel()
			return nil
		default:
			return countFSM
		}
	default:
		return pauseFSM
	}
}
