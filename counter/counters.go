package main

import "time"

type counter interface {
	display() []string
	addElapsed(time.Duration)
	addPause(time.Duration)
	duration() time.Duration
	finish(silent bool)
	finished() bool
	getRecords() CounterRecords
}

type nilCounter struct {
}

func (c *nilCounter) display() []string {
	return []string{}
}

func (c *nilCounter) addElapsed(gap time.Duration) {
}

func (c *nilCounter) addPause(gap time.Duration) {
}

func (c *nilCounter) duration() time.Duration {
	return 0
}

func (c *nilCounter) finish(silent bool) {
}

func (c *nilCounter) finished() bool {
	return true
}

func (c *nilCounter) getRecords() CounterRecords {
	return CounterRecords{}
}

// Counts down - like a timer
type downCounter struct {
	// Config
	dur          time.Duration
	silent       bool
	basicDisplay []string
	// Current State
	elapsed time.Duration
	paused  time.Duration
	// Recorded State
	times []TimeRecord
}

func newDownCounter(basicDisplay []string, durationStr string, silent bool) counter {
	dur, err := time.ParseDuration(durationStr)
	if err != nil {
		panic(err)
	}
	c := &downCounter{
		dur:    dur,
		silent: silent,
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

func (c *downCounter) duration() time.Duration {
	return c.dur
}

func (c *downCounter) remaining() time.Duration {
	return c.dur - c.elapsed
}

func (c *downCounter) display() []string {
	var d []string
	d = append(d, inSeconds(c.dur))
	d = append(d, inSeconds(c.remaining()))
	d = append(d, c.basicDisplay...)
	return d
}

func (c *downCounter) finished() bool {
	return c.elapsed > c.dur
}

func (c *downCounter) finish(silent bool) {
	if !silent && !c.silent {
		playSound()
	}
	if c.elapsed+c.paused > time.Second {
		time := TimeRecord{
			Elapsed: c.elapsed,
			Paused:  c.paused,
		}
		c.times = append(c.times, time)
	}
	c.elapsed = 0
	c.paused = 0
}

func (c *downCounter) getRecords() CounterRecords {
	times := make([]TimeRecord, len(c.times))
	copy(times, c.times)
	records := CounterRecords{
		Title: c.basicDisplay[0],
		Times: times,
	}
	return records
}
