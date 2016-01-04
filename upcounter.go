package main

import "time"

type upcounter struct {
	Title string
	counterData
}

func (c *upcounter) count() {
	runFSM(c)
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
