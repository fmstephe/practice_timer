package counters

import "time"

func Summarise(totalClock time.Duration, counters []Counter) *CountersSummary {
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
