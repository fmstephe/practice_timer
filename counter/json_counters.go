package main

import (
	"time"

	"github.com/fmstephe/countdown/tab"
)

type jsonCounter struct {
	Title    string
	Duration string
	Tab      string
}

func (c *jsonCounter) Generate(title, tabStr string) counter {
	var displayTab string
	if tabStr != "" {
		displayTab = tab.ExpandMotif(tabStr)
	}
	return newDownCounter([]string{title, displayTab}, c.Duration, false)
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
		genPause := cs.Pause.Generate("Up Next: "+c.Title, c.Tab)
		genCounter := c.Generate(c.Title, c.Tab)
		counters = append(counters, genPause)
		counters = append(counters, genCounter)
	}
	return counters
}

func (cs *multiCounters) summarise(totalClock time.Duration, counters []counter) *CountersSummary {
	records := make([]CounterRecords, 0)
	var totalElapsed time.Duration
	var totalPaused time.Duration
	for i, c := range counters {
		rs := c.getRecords()
		for _, t := range rs.Times {
			totalElapsed = totalElapsed + t.Elapsed
			totalPaused = totalPaused + t.Paused
		}
		if len(rs.Times) > 0 && i%2 != 0 {
			records = append(records, rs)
		}
	}
	summary := &CountersSummary{
		TotalElapsed: inSeconds(totalElapsed),
		TotalPaused:  inSeconds(totalPaused),
		TotalClock:   inSeconds(totalClock),
		Counters:     records,
	}
	return summary
}

type CountersSummary struct {
	Counters     []CounterRecords
	TotalClock   string
	TotalElapsed string
	TotalPaused  string
}

type CounterRecords struct {
	Title string
	Times []TimeRecord
}

type TimeRecord struct {
	Elapsed time.Duration
	Paused  time.Duration
}
