package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type titles struct {
	idx    int
	Titles []string
}

func (t *titles) next() string {
	title := t.Titles[t.idx%len(t.Titles)]
	t.idx++
	return title
}

func main() {
	generateRhythm()
}

func getTitles(fileName string) *titles {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	ts := &titles{}
	err = json.Unmarshal(bytes, ts)
	if err != nil {
		log.Fatal(err)
	}
	return ts
}

func applyTemplate(sessionData interface{}, templateFile, outFile string) string {
	bytes, err := ioutil.ReadFile(templateFile)
	if err != nil {
		log.Fatal(err)
	}
	t := template.Must(template.New("session").Parse(string(bytes)))
	fd, err := os.Create("../practice/" + outFile)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(fd, sessionData)
	if err != nil {
		log.Fatal(err)
	}
	return ""
}
