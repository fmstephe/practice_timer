package counters

import (
	"time"

	"github.com/fmstephe/countdown/tab"
)

type JsonCounter struct {
	Title    string
	Duration string
	Tab      string
}

func (c *JsonCounter) generate(title, tabStr string) Counter {
	var displayTab string
	if tabStr != "" {
		displayTab = tab.ExpandMotif(tabStr)
	}
	duration, err := time.ParseDuration(c.Duration)
	if err != nil {
		panic(err)
	}
	return NewDown([]string{title, displayTab}, duration, false)
}

type MultiCounters struct {
	Pause    JsonCounter
	Counters []JsonCounter
}

func (cs *MultiCounters) GenerateCounters() []Counter {
	var counters []Counter
	for _, c := range cs.Counters {
		genPause := cs.Pause.generate("Up Next: "+c.Title, c.Tab)
		genCounter := c.generate(c.Title, c.Tab)
		counters = append(counters, genPause)
		counters = append(counters, genCounter)
	}
	return counters
}
