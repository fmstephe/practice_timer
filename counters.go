package main

import "time"

// Some common mutable data for counters
type counterData struct {
	elapsed   time.Duration
	paused    time.Duration
	cancelled bool
	quiet     bool
}

func (d *counterData) restart() {
	d.elapsed = 0
	d.paused = 0
}

func (d *counterData) addPause(gap time.Duration) {
	d.paused += gap
}

func (d *counterData) addElapsed(gap time.Duration) {
	d.elapsed += gap
}

func (d *counterData) cancel() {
	d.cancelled = true
}

// Counts down - like a timer
type downCounter struct {
	title string
	total time.Duration
	counterData
}

func newDownCounter(title string, mins, secs int) counter {
	m := time.Duration(mins) * time.Minute
	s := time.Duration(secs) * time.Second
	total := m + s
	return &downCounter{
		title: title,
		total: total,
	}
}

func (c *downCounter) remaining() time.Duration {
	return c.total - c.elapsed + time.Second
}

func (c *downCounter) display() []string {
	return []string{c.title, inSeconds(c.elapsed), inSeconds(c.remaining())}
}

func (c *downCounter) finished() bool {
	return c.cancelled || c.elapsed > c.total
}

func (c *downCounter) finish() {
	if !c.cancelled && !c.quiet {
		playSound()
	}
}

// Counts up - like a stopwatch
type upCounter struct {
	title string
	counterData
}

func newUpCounter(title string) counter {
	return &upCounter{
		title: title,
	}
}

func (c *upCounter) display() []string {
	return []string{c.title, inSeconds(c.elapsed)}
}

func (c *upCounter) finished() bool {
	return c.cancelled
}

func (c *upCounter) finish() {
}
