package main

import "time"

type counter interface {
	display() []string
	isRest() bool
	addElapsed(time.Duration)
	addPause(time.Duration)
	finish(silent bool)
	finished() bool
	getRecords() CounterRecords
}

type nilCounter struct {
}

func (c *nilCounter) display() []string {
	return []string{}
}

func (c *nilCounter) isRest() bool {
	return false
}

func (c *nilCounter) addElapsed(gap time.Duration) {
}

func (c *nilCounter) addPause(gap time.Duration) {
}

func (c *nilCounter) finish(silent bool) {
}

func (c *nilCounter) finished() bool {
	return true
}

func (c *nilCounter) getRecords() CounterRecords {
	return nil
}

// Counts down - like a timer
type downCounter struct {
	// Config
	duration     time.Duration
	silent       bool
	rest         bool
	basicDisplay []string
	// Current State
	elapsed time.Duration
	paused  time.Duration
	// Recorded State
	records CounterRecords
}

func newDownCounter(basicDisplay []string, durationStr string, silent, rest bool) counter {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		panic(err)
	}
	c := &downCounter{
		duration: duration,
		silent:   silent,
		rest:     rest,
	}
	c.basicDisplay = basicDisplay
	return c
}

func (c *downCounter) addPause(gap time.Duration) {
	c.paused += gap
}

func (c *downCounter) addElapsed(gap time.Duration) {
	c.elapsed += gap
}

func (c *downCounter) remaining() time.Duration {
	return c.duration - c.elapsed + time.Second
}

func (c *downCounter) isRest() bool {
	return c.rest
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

func (c *downCounter) finish(silent bool) {
	if !silent && !c.silent {
		playSound()
	}
	if c.elapsed+c.paused > time.Second {
		record := &CounterRecord{
			Title:   c.basicDisplay[0],
			Elapsed: c.elapsed,
			Paused:  c.paused,
		}
		c.records = append(c.records, record)
	}
	c.elapsed = 0
	c.paused = 0
}

func (c *downCounter) getRecords() CounterRecords {
	records := make(CounterRecords, len(c.records))
	copy(records, c.records)
	return records
}
