package fsm

import (
	"strconv"
	"time"

	"github.com/fmstephe/countdown/counter/counters"
	"github.com/fmstephe/countdown/counter/stdin"
)

type fsmCounters struct {
	idx       int
	startTime time.Time
	counters  []counters.Counter
}

func (cs *fsmCounters) current() counters.Counter {
	if cs.idx >= len(cs.counters) {
		return counters.NewNilCounter()
	}
	return cs.counters[cs.idx]
}

func (cs *fsmCounters) display() []string {
	disp := []string{strconv.Itoa((cs.idx + 1)) + " of " + strconv.Itoa(len(cs.counters))}
	disp = append(disp, cs.remaining().String()+" remaining")
	disp = append(disp, cs.elapsedSeconds().String()+" Elapsed")
	disp = append(disp, cs.current().Display()...)
	return disp
}

func (cs *fsmCounters) elapsedSeconds() time.Duration {
	running := time.Now().Sub(cs.startTime)
	return running / time.Second * time.Second
}

func (cs *fsmCounters) remaining() time.Duration {
	var rem time.Duration
	for i := cs.idx; i < len(cs.counters); i++ {
		rem += cs.counters[i].Duration()
	}
	return rem
}

func (cs *fsmCounters) restart() {
	cs.current().Finish(true)
}

func (cs *fsmCounters) next() {
	cs.current().Finish(true)
	cs.idx++
	cs.restart()
}

func (cs *fsmCounters) prev() {
	cs.current().Finish(true)
	cs.idx--
	if cs.idx < 0 {
		cs.idx = 0
	}
	cs.restart()
}

func (cs *fsmCounters) addElapsed(gap time.Duration) {
	cs.current().AddElapsed(gap)
}

func (cs *fsmCounters) addPause(gap time.Duration) {
	cs.current().AddPause(gap)
}

func (cs *fsmCounters) quit() {
	cs.current().Finish(true)
	cs.idx = len(cs.counters)
}

func (cs *fsmCounters) finished() bool {
	if cs.idx >= len(cs.counters) {
		return true
	}
	if cs.current().Finished() {
		cs.current().Finish(false)
		cs.idx++
	}
	return cs.idx >= len(cs.counters)
}

func Run(counters []counters.Counter) {
	fsmC := &fsmCounters{
		startTime: time.Now(),
		counters:  counters,
	}
	tick := time.Now()
	f := countFSM
	for !fsmC.finished() {
		gap := time.Now().Sub(tick)
		tick = time.Now()
		f = f(fsmC, gap)
		time.Sleep(100 * time.Millisecond)
	}
	clearDisplay()
}

type counterFSM func(*fsmCounters, time.Duration) counterFSM

func countFSM(fsmC *fsmCounters, gap time.Duration) counterFSM {
	fsmC.addElapsed(gap)
	replaceText(fsmC.display())
	charChan := stdin.GetCharChan()
	select {
	case b := <-charChan:
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
	charChan := stdin.GetCharChan()
	select {
	case b := <-charChan:
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
