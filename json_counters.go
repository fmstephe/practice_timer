package main

import "log"

const (
	upMode   = "UP"
	downMode = "DOWN"
)

type jsonCounter struct {
	Mode     string
	Title    string
	Duration string
}

func (c *jsonCounter) Generate() counter {
	return c.GenerateTitled(c.Title)
}

func (c *jsonCounter) GenerateTitled(title string) counter {
	switch c.Mode {
	case upMode:
		return newUpCounter(title)
	case downMode:
		return newDownCounter(title, c.Duration)
	default:
		log.Fatalf("Bad mode: %+v", c)
		return nil
	}
}

type multiCounters struct {
	Pause    jsonCounter
	Counters []jsonCounter
}

func (cs multiCounters) countdown() []*CounterRecord {
	var counters []counter
	for _, c := range cs.Counters {
		genPause := cs.Pause.GenerateTitled("Up Next: " + c.Title)
		genCounter := c.Generate()
		counters = append(counters, genPause)
		counters = append(counters, genCounter)
	}
	var records []*CounterRecord
	for _, c := range counters {
		cRecord := runFSM(c)
		records = append(records, cRecord)
	}
	return records
}

type CounterRecord struct {
	Mode    string
	Title   string
	Elapsed string
	Paused  string
}
