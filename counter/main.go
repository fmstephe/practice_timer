package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/fmstephe/practice_timer/counter/counters"
	"github.com/fmstephe/practice_timer/counter/fsm"
)

var file = flag.String("f", "", "Optional path to a timer file")
var dow = flag.String("dow", "", "Optional path to a directory with a weeks worth of sessions")
var duration = flag.Duration("d", time.Minute+30*(time.Second), "The length of time for the countdown timer")
var repeats = flag.Int("r", 1, "The number of times to repeat the given duration")

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
	/// TODO the use of json everything here is very awkward
	multi := &counters.MultiCounters{}
	for i := 0; i < *repeats; i++ {
		jc := counters.JsonCounter{
			Title:    strconv.Itoa(i),
			Duration: (*duration).String(),
			Tab:      "",
		}
		multi.Counters = append(multi.Counters, jc)
	}
	runMulti(multi)
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
	runMulti(multi)
}

func runMulti(multi *counters.MultiCounters) {
	cs := multi.GenerateCounters()
	start := time.Now()
	fsm.Run(cs)
	wallClock := time.Now().Sub(start)
	summary := counters.Summarise(wallClock, cs)
	bytes, err := json.MarshalIndent(summary, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	println(string(bytes))
}

func todaysPractice(dirPath string) {
	day := time.Now().Weekday().String()
	fromFile(dirPath + "/" + day + ".json")
}
