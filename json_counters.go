package main

import (
	"log"
	"time"
)

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

func (cs *multiCounters) countdown() *CountersSummary {
	counters := cs.generateCounters()
	var records []*CounterRecord
	start := time.Now()
	for _, c := range counters {
		cRecord := runFSM(c)
		records = append(records, cRecord)
	}
	totalClock := time.Now().Sub(start)
	return cs.summarise(totalClock, records)
}

func (cs *multiCounters) generateCounters() []counter {
	var counters []counter
	for _, c := range cs.Counters {
		genPause := cs.Pause.GenerateTitled("Up Next: " + c.Title)
		genCounter := c.Generate()
		counters = append(counters, genPause)
		counters = append(counters, genCounter)
	}
	return counters
}

func (cs *multiCounters) summarise(totalClock time.Duration, records []*CounterRecord) *CountersSummary {
	summary := &CountersSummary{
		TotalClock: totalClock.String(),
		Counters:   records,
	}
	for _, r := range records {
		summary.TotalElapsed = addStringDurations(summary.TotalElapsed, r.Elapsed)
		summary.TotalPaused = addStringDurations(summary.TotalPaused, r.Paused)
	}
	return summary
}

func addStringDurations(ds1, ds2 string) string {
	d1, err := time.ParseDuration(ds1)
	if err != nil {
		log.Fatalf(err.Error())
	}
	d2, err := time.ParseDuration(ds2)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return (d1 + d2).String()
}

type CountersSummary struct {
	TotalClock   string
	TotalElapsed string
	TotalPaused  string
	Counters     []*CounterRecord
}

type CounterRecord struct {
	Mode    string
	Title   string
	Elapsed string
	Paused  string
}
