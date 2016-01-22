package main

import (
	"log"
	"time"

	"github.com/fmstephe/countdown/tab"
)

const (
	downMode  = "DOWN"
	pauseMode = "PAUSE"
)

type jsonCounter struct {
	Mode     string
	Title    string
	Duration string
	Tab      string
}

func (c *jsonCounter) Generate(basicDisplay []string) counter {
	switch c.Mode {
	case downMode:
		return newDownCounter(basicDisplay, c.Duration, false)
	case pauseMode:
		basicDisplay[0] = "Up Next: " + basicDisplay[0]
		return newDownCounter(basicDisplay, c.Duration, true)
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
	fsmC := &fsmCounters{
		counters: counters,
	}
	start := time.Now()
	runFSM(fsmC)
	totalClock := time.Now().Sub(start)
	return cs.summarise(totalClock, counters)
}

func (cs *multiCounters) generateCounters() []counter {
	var counters []counter
	for _, c := range cs.Counters {
		if c.Tab == "" {
			genPause := cs.Pause.Generate([]string{c.Title})
			genCounter := c.Generate([]string{c.Title})
			counters = append(counters, genPause)
			counters = append(counters, genCounter)
		} else {
			genPause := cs.Pause.Generate([]string{c.Title, tab.ExpandMotif(c.Tab)})
			genCounter := c.Generate([]string{c.Title, tab.ExpandMotif(c.Tab)})
			counters = append(counters, genPause)
			counters = append(counters, genCounter)
		}
	}
	return counters
}

func (cs *multiCounters) summarise(totalClock time.Duration, counters []counter) *CountersSummary {
	var records []*CounterRecord
	var totalElapsed time.Duration
	var totalPaused time.Duration
	for _, c := range counters {
		r := c.getRecord()
		totalElapsed = totalElapsed + r.Elapsed
		totalPaused = totalPaused + r.Paused
		records = append(records, r)
	}
	summary := &CountersSummary{
		TotalElapsed: totalElapsed,
		TotalPaused:  totalPaused,
		TotalClock:   totalClock,
		Counters:     records,
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
	Display []string
	Elapsed time.Duration
	Paused  time.Duration
}
