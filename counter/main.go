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
var rod = flag.Bool("rod", false, "If set the counter will select the rhythm practice file for the day of the week")
var lod = flag.Bool("lod", false, "If set the counter will select the lead practice file for the day of the week")

func main() {
	flag.Parse()
	switch {
	case *rod:
		todaysPractice("rhythm")
	case *lod:
		todaysPractice("lead")
	case *file != "":
		fromFile(*file)
	default:
		simple()
	}
}

func simple() {
	c := newDownCounter([]string{*title}, *duration, false)
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

func todaysPractice(category string) {
	day := time.Now().Weekday().String()
	fromFile("../practice/" + "/" + category + "/" + day + ".json")
}
