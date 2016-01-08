package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

var file = flag.String("f", "", "Optional path to a timer file")
var title = flag.String("h", "", "An optional title to display above timer")
var duration = flag.String("d", "1m30s", "The length of time for the countdown timer")

func main() {
	flag.Parse()
	if *file == "" {
		simple()
	} else {
		fromFile()
	}
}

func simple() {
	c := newDownCounter(*title, *duration)
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
	records := counters.countdown()
	bytes, err = json.Marshal(records)
	if err != nil {
		log.Fatal(err)
	}
	clearDisplay()
	println(string(bytes))

}
