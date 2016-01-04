package main

import "time"

type downcounter struct {
	Title   string
	Minutes int
	Seconds int
	counterData
}

func (c *downcounter) count(quiet bool) {
	runFSM(c)
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
