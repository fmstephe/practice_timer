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

type sessionData struct {
	OneChordRhythm1 string
	OneChordRhythm2 string
	SingleNote1     string
	//
	OneChordRhythm3 string
	OneChordRhythm4 string
	SingleNote2     string
	//
	TwoChordRhythm1 string
	TwoChordRhythm2 string
	SingleNote3     string
	//
	Barre1      string
	NeckRhythm1 string
	SingleNote4 string
	SingleNote5 string
	//
	Barre2      string
	NeckRhythm2 string
	SingleNote6 string
	SingleNote7 string
	//
	Barre3      string
	NeckRhythm3 string
	SingleNote8 string
	SingleNote9 string
}

func main() {
	barre := getTitles("barre.json")
	oneChordRhythm := getTitles("oneChordRhythm.json")
	twoChordRhythm := getTitles("twoChordRhythm.json")
	neckRhythm := getTitles("neckRhythm.json")
	singleNote := getTitles("singleNote.json")
	sds := make([]sessionData, 7)
	for i := range sds {
		sds[i].OneChordRhythm1 = oneChordRhythm.next()
		sds[i].OneChordRhythm2 = oneChordRhythm.next()
		sds[i].SingleNote1 = singleNote.next()
		//
		sds[i].OneChordRhythm3 = oneChordRhythm.next()
		sds[i].OneChordRhythm4 = oneChordRhythm.next()
		sds[i].SingleNote2 = singleNote.next()
		//
		sds[i].TwoChordRhythm1 = twoChordRhythm.next()
		sds[i].TwoChordRhythm2 = twoChordRhythm.next()
		sds[i].SingleNote3 = singleNote.next()
		//
		sds[i].Barre1 = barre.next()
		sds[i].NeckRhythm1 = neckRhythm.next()
		sds[i].SingleNote4 = singleNote.next()
		sds[i].SingleNote5 = singleNote.next()
		//
		sds[i].Barre2 = barre.next()
		sds[i].NeckRhythm2 = neckRhythm.next()
		sds[i].SingleNote6 = singleNote.next()
		sds[i].SingleNote7 = singleNote.next()
		//
		sds[i].Barre3 = barre.next()
		sds[i].NeckRhythm3 = neckRhythm.next()
		sds[i].SingleNote8 = singleNote.next()
		sds[i].SingleNote9 = singleNote.next()
	}
	applyTemplate(sds[0], "Monday.json")
	applyTemplate(sds[1], "Tuesday.json")
	applyTemplate(sds[2], "Wednesday.json")
	applyTemplate(sds[3], "Thursday.json")
	applyTemplate(sds[4], "Friday.json")
	applyTemplate(sds[5], "Saturday.json")
	applyTemplate(sds[6], "Sunday.json")
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

func applyTemplate(sd sessionData, fileName string) string {
	bytes, err := ioutil.ReadFile("Template.json")
	if err != nil {
		log.Fatal(err)
	}
	t := template.Must(template.New("session").Parse(string(bytes)))
	fd, err := os.Create("../practice/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(fd, sd)
	if err != nil {
		log.Fatal(err)
	}
	return ""
}
