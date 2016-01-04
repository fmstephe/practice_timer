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
type downcounter struct {
	Title   string
	Minutes int
	Seconds int
	counterData
}

func newDowncounter(title string, mins, secs int) *downcounter {
	return &downcounter{
		Title:   title,
		Minutes: mins,
		Seconds: secs,
	}
}

func (c *downcounter) total() time.Duration {
	m := time.Duration(c.Minutes) * time.Minute
	s := time.Duration(c.Seconds) * time.Second
	return m + s
}

func (c *downcounter) remaining() time.Duration {
	return c.total() - c.elapsed() + time.Second
}

func (c *downcounter) display() []string {
	return []string{c.Title, inSeconds(c.elapsed()), inSeconds(c.remaining())}
}

func (c *downcounter) finished() bool {
	return c.elapsed() > c.total()
}

func (c *downcounter) finish() {
	if !c.quiet {
		playSound()
	}
}

// Counts up - like a stopwatch
type upcounter struct {
	Title string
	counterData
}

func newUpcounter(title string) counter {
	return &upcounter{
		Title: title,
	}
}

func (c *upcounter) elapsed() time.Duration {
	return time.Now().Sub(c.start)
}

func (c *upcounter) display() []string {
	return []string{c.Title, inSeconds(c.elapsed())}
}

func (c *upcounter) finished() bool {
	return false
}

func (c *upcounter) finish() {
}
