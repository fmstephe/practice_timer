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

func runFSM(c counter) (*CounterRecord, bool) {
	c.restart()
	tick := time.Now()
	f := countFSM
	var quit bool
	for !c.finished() {
		gap := time.Now().Sub(tick)
		tick = time.Now()
		f, quit = f(c, gap)
		if quit {
			return c.finish(), true
		}
		time.Sleep(100 * time.Millisecond)
	}
	return c.finish(), false
}

type counterFSM func(counter, time.Duration) (counterFSM, bool)

func countFSM(c counter, gap time.Duration) (counterFSM, bool) {
	c.addElapsed(gap)
	replaceText(c.display())
	select {
	case b := <-stdinChars:
		switch b {
		case spaceChar:
			return pauseFSM, false
		case qChar:
			c.cancel()
			return nil, true
		case rChar:
			c.restart()
			return countFSM, false
		case cChar:
			c.cancel()
			return nil, false
		default:
			return countFSM, false
		}
	default:
		return countFSM, false
	}
}

func pauseFSM(c counter, gap time.Duration) (counterFSM, bool) {
	c.addPause(gap)
	replaceText(c.display(), "PAUSED")
	select {
	case b := <-stdinChars:
		switch b {
		case spaceChar:
			return countFSM, false
		case qChar:
			c.cancel()
			return nil, true
		case rChar:
			c.restart()
			return pauseFSM, false
		case cChar:
			c.cancel()
			return nil, false
		default:
			return countFSM, false
		}
	default:
		return pauseFSM, false
	}
}
