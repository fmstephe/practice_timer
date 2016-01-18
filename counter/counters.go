package main

import "time"

type counter interface {
	display() []string
	restart()
	addElapsed(time.Duration)
	addPause(time.Duration)
	finish()
	finished() bool
	getRecord() *CounterRecord
}

type nilCounter struct {
}

func (c *nilCounter) display() []string {
	return []string{}
}

func (c *nilCounter) restart() {
}

func (c *nilCounter) addElapsed(gap time.Duration) {
}

func (c *nilCounter) addPause(gap time.Duration) {
}

func (c *nilCounter) finish() {
}

func (c *nilCounter) finished() bool {
	return true
}

func (c *nilCounter) getRecord() *CounterRecord {
	return nil
}

// Some common mutable data for counters
type counterData struct {
	elapsed time.Duration
	paused  time.Duration
	silent  bool
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

// Counts down - like a timer
type downCounter struct {
	title    string
	duration time.Duration
	silent   bool
	counterData
}

func newDownCounter(title, durationStr string, silent bool) counter {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		panic(err)
	}
	return &downCounter{
		title:    title,
		duration: duration,
		silent:   silent,
	}
}

func (c *downCounter) remaining() time.Duration {
	return c.duration - c.elapsed + time.Second
}

func (c *downCounter) display() []string {
	return []string{c.title, inSeconds(c.elapsed), inSeconds(c.remaining())}
}

func (c *downCounter) finished() bool {
	return c.elapsed > c.duration
}

func (c *downCounter) finish() {
	if !c.silent {
		playSound()
	}
}

func (c *downCounter) getRecord() *CounterRecord {
	return &CounterRecord{
		Mode:    downMode,
		Title:   c.title,
		Elapsed: c.elapsed,
		Paused:  c.paused,
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
	return false
}

func (c *upCounter) finish() {
}

func (c *upCounter) getRecord() *CounterRecord {
	return &CounterRecord{
		Mode:    upMode,
		Title:   c.title,
		Elapsed: c.elapsed,
		Paused:  c.paused,
	}
}
