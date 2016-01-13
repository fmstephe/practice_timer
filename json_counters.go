package main

import (
	"log"
	"time"
)

var (
	zeroDuration time.Duration
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
	return c.GenerateTitled(c.Title, false)
}

func (c *jsonCounter) GenerateTitled(title string, silent bool) counter {
	switch c.Mode {
	case upMode:
		return newUpCounter(title)
	case downMode:
		return newDownCounter(title, c.Duration, silent)
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
		genPause := cs.Pause.GenerateTitled("Up Next: "+c.Title, true)
		genCounter := c.Generate()
		counters = append(counters, genPause)
		counters = append(counters, genCounter)
	}
	return counters
}

func (cs *multiCounters) summarise(totalClock time.Duration, records []*CounterRecord) *CountersSummary {
	summary := &CountersSummary{
		TotalElapsed: zeroDuration,
		TotalPaused:  zeroDuration,
		TotalClock:   totalClock,
		Counters:     records,
	}
	for _, r := range records {
		summary.TotalElapsed = summary.TotalElapsed + r.Elapsed
		summary.TotalPaused = summary.TotalPaused + r.Paused
	}
	return summary
}

type CountersSummary struct {
	TotalClock   time.Duration
	TotalElapsed time.Duration
	TotalPaused  time.Duration
	Counters     []*CounterRecord
}

type CounterRecord struct {
	Mode    string
	Title   string
	Elapsed time.Duration
	Paused  time.Duration
}
