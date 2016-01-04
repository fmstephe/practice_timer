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

func main() {
	flag.Parse()
	if *file == "" {
		simple()
	} else {
		fromFile()
	}
}

func simple() {
	c := &downcounter{
		Title:   *title,
		Minutes: *minutes,
		Seconds: *seconds,
	}
	runFSM(c)
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
	counters.countdown()
}
