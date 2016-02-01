package main

import (
	"time"

	"github.com/fmstephe/countdown/tab"
)

type jsonCounter struct {
	IsRest   bool
	Title    string
	Duration string
	Tab      string
}

func (c *jsonCounter) Generate(title, tabStr string) counter {
	var displayTab string
	if tabStr != "" {
		displayTab = tab.ExpandMotif(tabStr)
	}
	if c.IsRest {
		title = "Up Next: " + title
	}
	return newDownCounter([]string{title, displayTab}, c.Duration, false, c.IsRest)
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
		genPause := cs.Pause.Generate(c.Title, c.Tab)
		genCounter := c.Generate(c.Title, c.Tab)
		counters = append(counters, genPause)
		counters = append(counters, genCounter)
	}
	return counters
}

func (cs *multiCounters) summarise(totalClock time.Duration, counters []counter) *CountersSummary {
	var records []CounterRecords
	var totalElapsed time.Duration
	var totalPaused time.Duration
	for _, c := range counters {
		if c.isRest() {
			continue
		}
		rs := c.getRecords()
		for _, r := range rs {
			totalElapsed = totalElapsed + r.Elapsed
			totalPaused = totalPaused + r.Paused
		}
		if len(rs) > 0 {
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
	TotalClock   string
	TotalElapsed string
	TotalPaused  string
	Counters     []CounterRecords
}

type CounterRecords []*CounterRecord

type CounterRecord struct {
	Title   string
	Elapsed time.Duration
	Paused  time.Duration
}
