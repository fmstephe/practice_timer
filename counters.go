package main

type multiCounters struct {
	Pause    downcounter
	Counters []downcounter
}

func (cs multiCounters) countdown(quiet bool) {
	for _, c := range cs.Counters {
		pause := cs.Pause
		pause.Title = "Up Next: " + c.Title
		pause.count(true)
		c.count(quiet)
	}
}
