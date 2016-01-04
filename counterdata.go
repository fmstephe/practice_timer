package main

import "time"

type counterData struct {
	// Separate these mutable fields
	start      time.Time
	pauseStart time.Time
	pauses     []time.Duration
}

func (d *counterData) totalPaused() time.Duration {
	var paused time.Duration
	for _, p := range d.pauses {
		paused += p
	}
	return paused
}

func (d *counterData) pauseElapsed() time.Duration {
	return time.Now().Sub(d.pauseStart)
}

func (d *counterData) restart() {
	d.start = time.Now()
	d.pauseStart = d.start
	d.pauses = nil
}
