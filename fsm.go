package main

import "time"

type fsmCounters struct {
	idx      int
	counters []counter
}

func (cs *fsmCounters) current() counter {
	return cs.counters[cs.idx]
}

func (cs *fsmCounters) display() []string {
	return cs.current().display()
}

func (cs *fsmCounters) restart() {
	cs.current().restart()
}

func (cs *fsmCounters) addElapsed(gap time.Duration) {
	cs.current().addElapsed(gap)
}

func (cs *fsmCounters) addPause(gap time.Duration) {
	cs.current().addPause(gap)
}

func (cs *fsmCounters) cancel() {
	cs.current().cancel()
}

func (cs *fsmCounters) quit() {
	cs.current().cancel()
	cs.idx = len(cs.counters)
}

func (cs *fsmCounters) finished() bool {
	if cs.idx == len(cs.counters) {
		return true
	}
	if cs.current().finished() {
		cs.current().finish()
		cs.idx++
	}
	return cs.idx == len(cs.counters)
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
		case cChar:
			fsmC.cancel()
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
		case cChar:
			fsmC.cancel()
			return pauseFSM
		default:
			return pauseFSM
		}
	default:
		return pauseFSM
	}
}
