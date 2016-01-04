package main

import "log"

const (
	upMode   = "UP"
	downMode = "DOWN"
)

type JsonCounter struct {
	Mode    string
	Title   string
	Minutes int
	Seconds int
}

func (c *JsonCounter) GenerateCounter() counter {
	switch c.Mode {
	case upMode:
		return newUpCounter(c.Title)
	case downMode:
		return newDowncounter(c.Title, c.Minutes, c.Seconds)
	default:
		log.Fatalf("Bad mode: %+v", c)
		return nil
	}
}

type multiCounters struct {
	Pause    downcounter
	Counters []downcounter
}

func (cs multiCounters) countdown() {
	for _, c := range cs.Counters {
		pause := cs.Pause
		pause.Title = "Up Next: " + c.Title
		runFSM(&pause)
		runFSM(&c)
	}
}
