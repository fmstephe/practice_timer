package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

var file = flag.String("f", "", "Optional path to a timer file")
var title = flag.String("h", "", "An optional title to display above timer")
var minutes = flag.Int("m", 0, "The number of minutes for the timer to run")
var seconds = flag.Int("s", 90, "The number of seconds for the timer to run")
var quiet = flag.Bool("q", false, "If true the applause will not play at end of timer")

func main() {
	flag.Parse()
	if *file == "" {
		simple()
	} else {
		fromFile()
	}
}

func simple() {
	c := &counter{
		Title:   *title,
		Minutes: *minutes,
		Seconds: *seconds,
	}
	c.countdown(*quiet)
}

func fromFile() {
	bytes, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}
	counters := &multiCounters{}
	err = json.Unmarshal(bytes, counters)
	if err != nil {
		log.Fatal(err)
	}
	counters.countdown(*quiet)
}
