package main

import (
	"strconv"
	"time"
)

type fsmCounters struct {
	idx      int
	counters []counter
}

func (cs *fsmCounters) current() counter {
	if cs.idx >= len(cs.counters) {
		return &nilCounter{}
	}
	return cs.counters[cs.idx]
}

func (cs *fsmCounters) display() []string {
	disp := []string{strconv.Itoa((cs.idx+2)/2) + " of " + strconv.Itoa(len(cs.counters)/2)}
	disp = append(disp, cs.current().display()...)
	return disp
}

func (cs *fsmCounters) restart() {
	cs.current().restart()
}

func (cs *fsmCounters) next() {
	cs.idx++
	cs.restart()
}

func (cs *fsmCounters) prev() {
	cs.idx--
	if cs.idx < 0 {
		cs.idx = 0
	}
	cs.restart()
}

func (cs *fsmCounters) addElapsed(gap time.Duration) {
	cs.current().addElapsed(gap)
}

func (cs *fsmCounters) addPause(gap time.Duration) {
	cs.current().addPause(gap)
}

func (cs *fsmCounters) quit() {
	cs.idx = len(cs.counters)
}

func (cs *fsmCounters) finished() bool {
	if cs.idx >= len(cs.counters) {
		return true
	}
	if cs.current().finished() {
		cs.current().finish()
		cs.idx++
	}
	return cs.idx >= len(cs.counters)
}

func runFSM(fsmC *fsmCounters) {
	tick := time.Now()
	f := countFSM
	for !fsmC.finished() {
		gap := time.Now().Sub(tick)
		tick = time.Now()
		f = f(fsmC, gap)
		time.Sleep(100 * time.Millisecond)
	}
}

type counterFSM func(*fsmCounters, time.Duration) counterFSM

func countFSM(fsmC *fsmCounters, gap time.Duration) counterFSM {
	fsmC.addElapsed(gap)
	replaceText(fsmC.display())
	select {
	case b := <-stdinChars:
		switch b {
		case spaceChar:
			return pauseFSM
		case qChar:
			fsmC.quit()
			return nil
		case rChar:
			fsmC.restart()
			return countFSM
		case aChar:
			fsmC.prev()
			return countFSM
		case dChar:
			fsmC.next()
			return countFSM
		default:
			return countFSM
		}
	default:
		return countFSM
	}
}

func pauseFSM(fsmC *fsmCounters, gap time.Duration) counterFSM {
	fsmC.addPause(gap)
	replaceText(fsmC.display(), "PAUSED")
	select {
	case b := <-stdinChars:
		switch b {
		case spaceChar:
			return countFSM
		case qChar:
			fsmC.quit()
			return nil
		case rChar:
			fsmC.restart()
			return pauseFSM
		case aChar:
			fsmC.prev()
			return pauseFSM
		case dChar:
			fsmC.next()
			return pauseFSM
		default:
			return pauseFSM
		}
	default:
		return pauseFSM
	}
}
