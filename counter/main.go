package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"time"
)

var file = flag.String("f", "", "Optional path to a timer file")
var dow = flag.String("dow", "", "Optional path to a directory with a weeks worth of sessions")
var title = flag.String("h", "", "An optional title to display above timer")
var duration = flag.String("d", "1m30s", "The length of time for the countdown timer")

func main() {
	flag.Parse()
	switch {
	case *dow != "":
		todaysPractice(*dow)
	case *file != "":
		fromFile(*file)
	default:
		simple()
	}
}

func simple() {
	c := newDownCounter([]string{*title}, *duration, false, false)
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
	bytes, err = json.MarshalIndent(records, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	clearDisplay()
	println(string(bytes))
}

func todaysPractice(dirPath string) {
	day := time.Now().Weekday().String()
	fromFile(dirPath + "/" + day + ".json")
}
