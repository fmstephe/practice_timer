package main

import "log"

const (
	upMode   = "UP"
	downMode = "DOWN"
)

type jsonCounter struct {
	Mode    string
	Title   string
	Minutes int
	Seconds int
}

func (c *jsonCounter) Generate() counter {
	return c.GenerateTitled(c.Title)
}

func (c *jsonCounter) GenerateTitled(title string) counter {
	switch c.Mode {
	case upMode:
		return newUpCounter(title)
	case downMode:
		return newDownCounter(title, c.Minutes, c.Seconds)
	default:
		log.Fatalf("Bad mode: %+v", c)
		return nil
	}
}

type multiCounters struct {
	Pause    jsonCounter
	Counters []jsonCounter
}

func (cs multiCounters) countdown() {
	for _, c := range cs.Counters {
		genPause := cs.Pause.GenerateTitled("Up Next: " + c.Title)
		genCounter := c.Generate()
		runFSM(genPause)
		runFSM(genCounter)
	}
}

type CounterRecord struct {
	Mode          string
	Title         string
	Minutes       int
	Seconds       int
	PausedMinutes int
	PausedSeconds int
}
