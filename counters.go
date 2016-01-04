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
	Title   string
	Minutes int
	Seconds int
	counterData
}

func newDownCounter(title string, mins, secs int) *downCounter {
	return &downCounter{
		Title:   title,
		Minutes: mins,
		Seconds: secs,
	}
}

func (c *downCounter) total() time.Duration {
	m := time.Duration(c.Minutes) * time.Minute
	s := time.Duration(c.Seconds) * time.Second
	return m + s
}

func (c *downCounter) remaining() time.Duration {
	return c.total() - c.elapsed() + time.Second
}

func (c *downCounter) display() []string {
	return []string{c.Title, inSeconds(c.elapsed()), inSeconds(c.remaining())}
}

func (c *downCounter) finished() bool {
	return c.elapsed() > c.total()
}

func (c *downCounter) finish() {
	if !c.quiet {
		playSound()
	}
}

// Counts up - like a stopwatch
type upCounter struct {
	Title string
	counterData
}

func newUpCounter(title string) counter {
	return &upCounter{
		Title: title,
	}
}

func (c *upCounter) elapsed() time.Duration {
	return time.Now().Sub(c.start)
}

func (c *upCounter) display() []string {
	return []string{c.Title, inSeconds(c.elapsed())}
}

func (c *upCounter) finished() bool {
	return false
}

func (c *upCounter) finish() {
}
