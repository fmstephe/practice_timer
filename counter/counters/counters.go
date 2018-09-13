package counters

import (
	"log"
	"os/exec"
	"time"
)

type Counter interface {
	Display() []string
	AddElapsed(time.Duration)
	AddPause(time.Duration)
	Duration() time.Duration
	Finish(silent bool)
	Finished() bool
	getRecords() CounterRecords
}

type nilCounter struct {
}

func NewNilCounter() Counter {
	return &nilCounter{}
}

func (c *nilCounter) Display() []string {
	return []string{}
}

func (c *nilCounter) AddElapsed(gap time.Duration) {
}

func (c *nilCounter) AddPause(gap time.Duration) {
}

func (c *nilCounter) Duration() time.Duration {
	return 0
}

func (c *nilCounter) Finish(silent bool) {
}

func (c *nilCounter) Finished() bool {
	return true
}

func (c *nilCounter) getRecords() CounterRecords {
	return CounterRecords{}
}

// Counts down - like a timer
type downCounter struct {
	// Config
	duration     time.Duration
	silent       bool
	basicDisplay []string
	// Current State
	elapsed time.Duration
	paused  time.Duration
	// Recorded State
	times []TimeRecord
}

func NewDown(basicDisplay []string, duration time.Duration, silent bool) Counter {
	c := &downCounter{
		duration: duration,
		silent:   silent,
	}
	c.basicDisplay = basicDisplay
	return c
}

func (c *downCounter) AddPause(gap time.Duration) {
	c.paused += gap
}

func (c *downCounter) AddElapsed(gap time.Duration) {
	c.elapsed += gap
}

func (c *downCounter) Duration() time.Duration {
	return c.duration
}

func (c *downCounter) Display() []string {
	var d []string
	d = append(d, inSeconds(c.remaining()))
	d = append(d, c.basicDisplay...)
	d = append(d, c.progressBoxes())
	return d
}

func (c *downCounter) progressBoxes() string {
	unitSize := int(c.duration) / 80
	boxes := ""
	count := 0
	for i := 0; i < int(c.remaining()); i += unitSize {
		boxes += "█"
		count++
	}
	for i := count; i < 79; i++ {
		boxes += " "
	}
	boxes += "█"
	return boxes
}

func (c *downCounter) Finished() bool {
	return c.elapsed > c.duration
}

func (c *downCounter) Finish(silent bool) {
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

func (c *downCounter) remaining() time.Duration {
	return c.duration - c.elapsed
}

func inSeconds(d time.Duration) string {
	d = ((d / time.Second) * time.Second)
	if d == 0 {
		return "0s"
	}
	return ((d / time.Second) * time.Second).String()
}

func playSound() {
	cmd := exec.Command("afplay", "-v", "0.05", "clap.wav")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
