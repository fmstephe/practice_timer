package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/fmstephe/practice_timer/counter/counters"
	"github.com/fmstephe/practice_timer/counter/fsm"
)

var fileFlag = flag.String("f", "", "Optional path to a timer file")
var dowFlag = flag.String("dow", "", "Optional path to a directory with a weeks worth of sessions")
var durationsFlag = flag.String("d", "15s,2m", "Comma separated list of timer durations")
var repeatsFlag = flag.Int("r", 1, "The number of times to repeat the given duration")

var durations []time.Duration

func main() {
	flag.Parse()
	if err := parseDurations(); err != nil {
		fmt.Printf("Error parsing timer durations %s - %s", *durationsFlag, err)
	}

	switch {
	case *dowFlag != "":
		todaysPractice(*dowFlag)
	case *fileFlag != "":
		fromFile(*fileFlag)
	default:
		simple()
	}
}

func parseDurations() error {
	durationStrs := strings.Split(*durationsFlag, ",")
	for _, durationStr := range durationStrs {
		duration, err := time.ParseDuration(durationStr)
		if err != nil {
			return err
		}
		durations = append(durations, duration)
	}
	return nil
}

func simple() {
	/// TODO the use of json everything here is very awkward
	multi := &counters.MultiCounters{}
	for i := 0; i < *repeatsFlag; i++ {
		for _, dur := range durations {
			jc := counters.JsonCounter{
				Title:    strconv.Itoa(i),
				Duration: dur.String(),
				Tab:      "",
			}
			multi.Counters = append(multi.Counters, jc)
		}
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
