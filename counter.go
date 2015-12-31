package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

type multiCounters struct {
	Pause    counter
	Counters []counter
}

func (cs multiCounters) countdown(quiet bool) {
	for _, c := range cs.Counters {
		pause := cs.Pause
		pause.Title = "Up Next: " + c.Title
		pause.countdown(true)
		c.countdown(quiet)
	}
}

type counter struct {
	Title   string
	Minutes int
	Seconds int
}

func (c counter) countdown(quiet bool) {
	start := time.Now()
	for !c.updateDisplay(start) {
		time.Sleep(time.Second)
	}
	c.completeDisplay()
	if !quiet {
		c.playSound()
	}
}

func (c counter) updateDisplay(start time.Time) bool {
	elapsed := time.Now().Sub(start)
	remainder := c.Total() - elapsed + time.Second
	clearDisplay()
	if c.Title != "" {
		println(c.Title)
	}
	println(inSeconds(elapsed))
	println(inSeconds(remainder))
	return elapsed >= c.Total()
}

func (c counter) completeDisplay() {
	clearDisplay()
	if c.Title != "" {
		println(c.Title)
	}
	println(inSeconds(c.Total()))
	println(inSeconds(c.Total()))
}

func (c counter) Total() time.Duration {
	m := time.Duration(c.Minutes) * time.Minute
	s := time.Duration(c.Seconds) * time.Second
	return m + s
}

func clearDisplay() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func inSeconds(d time.Duration) string {
	if d == 0 {
		return "0s"
	}
	return ((d / time.Second) * time.Second).String()
}

func (c counter) playSound() {
	clearDisplay()
	cmd := exec.Command("paplay", "clap.wav")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
