package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
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
	generateLead()
}

func getTitles(fileName string) *titles {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(bytes), "\n")
	ts := &titles{
		Titles: removeEmpty(lines),
	}
	return ts
}

func removeEmpty(ss []string) []string {
	nss := make([]string, 0, len(ss))
	for _, s := range ss {
		if s != "" {
			nss = append(nss, s)
		}
	}
	return nss
}

func applyTemplate(sessionData interface{}, category string, dow time.Weekday) string {
	bytes, err := ioutil.ReadFile(category + "/template.json")
	if err != nil {
		log.Fatal(err)
	}
	t := template.Must(template.New("session").Parse(string(bytes)))
	err = os.MkdirAll("../practice/"+category+"/", 0777)
	if err != nil {
		log.Fatal(err)
	}
	fd, err := os.Create("../practice/" + category + "/" + dow.String() + ".json")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(fd, sessionData)
	if err != nil {
		log.Fatal(err)
	}
	return ""
}
