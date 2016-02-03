package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"time"

	"github.com/fmstephe/countdown/counter/counters"
	"github.com/fmstephe/countdown/counter/fsm"
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
	c := counters.NewDown([]string{*title}, *duration, false)
	fsm.Run([]counters.Counter{c})
}

func fromFile(fileName string) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	multi := &counters.MultiCounters{}
	err = json.Unmarshal(bytes, multi)
	if err != nil {
		log.Fatal(err)
	}
	cs := multi.GenerateCounters()
	start := time.Now()
	fsm.Run(cs)
	wallClock := time.Now().Sub(start)
	summary := counters.Summarise(wallClock, cs)
	bytes, err = json.MarshalIndent(summary, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	println(string(bytes))
}

func todaysPractice(dirPath string) {
	day := time.Now().Weekday().String()
	fromFile(dirPath + "/" + day + ".json")
}
