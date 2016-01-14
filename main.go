package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"time"
)

var file = flag.String("f", "", "Optional path to a timer file")
var title = flag.String("h", "", "An optional title to display above timer")
var duration = flag.String("d", "1m30s", "The length of time for the countdown timer")
var pod = flag.Bool("pod", false, "If set the counter will select the practice file for the day of the week")

func main() {
	flag.Parse()
	switch {
	case *pod:
		practiceOfTheDay()
	case *file != "":
		fromFile(*file)
	default:
		simple()
	}
}

func simple() {
	c := newDownCounter(*title, *duration, false)
	fsmC := &fsmCounters{
		counters: []counter{c},
	}
	runFSM(fsmC)
}

func fromFile(fileName string) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	counters := &multiCounters{}
	err = json.Unmarshal(bytes, counters)
	if err != nil {
		log.Fatal(err)
	}
	records := counters.countdown()
	bytes, err = json.Marshal(records)
	if err != nil {
		log.Fatal(err)
	}
	clearDisplay()
	println(string(bytes))
}

func practiceOfTheDay() {
	day := time.Now().Weekday().String()
	fromFile("practice/" + day + ".json")
}
