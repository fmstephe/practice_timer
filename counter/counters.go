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
	elapsed      time.Duration
	paused       time.Duration
	basicDisplay []string
	silent       bool
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
	duration time.Duration
	silent   bool
	counterData
}

func newDownCounter(basicDisplay []string, durationStr string, silent bool) counter {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		panic(err)
	}
	c := &downCounter{
		duration: duration,
		silent:   silent,
	}
	c.basicDisplay = basicDisplay
	return c
}

func (c *downCounter) remaining() time.Duration {
	return c.duration - c.elapsed + time.Second
}

func (c *downCounter) display() []string {
	var d []string
	d = append(d, inSeconds(c.duration))
	d = append(d, inSeconds(c.remaining()))
	d = append(d, c.basicDisplay...)
	return d
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
		Display: c.basicDisplay,
		Elapsed: c.elapsed,
		Paused:  c.paused,
	}
}
