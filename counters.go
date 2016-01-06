package main

import "time"

// Some common mutable data for counters
type counterData struct {
	start  time.Time
	paused time.Duration
	quiet  bool
}

func (d *counterData) elapsed() time.Duration {
	return time.Now().Sub(d.start) - d.paused
}

func (d *counterData) restart() {
	d.start = time.Now()
	d.paused = 0
}

func (d *counterData) addPause(gap time.Duration) {
	d.paused += gap
}

// Counts down - like a timer
type downCounter struct {
	title string
	total time.Duration
	counterData
}

func newDownCounter(title string, mins, secs int) *downCounter {
	m := time.Duration(mins) * time.Minute
	s := time.Duration(secs) * time.Second
	total := m + s
	return &downCounter{
		title: title,
		total: total,
	}
}

func (c *downCounter) remaining() time.Duration {
	return c.total - c.elapsed() + time.Second
}

func (c *downCounter) display() []string {
	return []string{c.title, inSeconds(c.elapsed()), inSeconds(c.remaining())}
}

func (c *downCounter) finished() bool {
	return c.elapsed() > c.total
}

func (c *downCounter) finish() {
	if !c.quiet {
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

func (c *upCounter) elapsed() time.Duration {
	return time.Now().Sub(c.start)
}

func (c *upCounter) display() []string {
	return []string{c.title, inSeconds(c.elapsed())}
}

func (c *upCounter) finished() bool {
	return false
}

func (c *upCounter) finish() {
}
