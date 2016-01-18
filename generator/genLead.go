package main

import (
	"time"
)

type leadSession struct {
	Riff1         string
	Riff2         string
	Riff3         string
	Riff4         string
	Classical1    string
	Classical2    string
	Classical3    string
	Classical4    string
	Classical5    string
	Classical6    string
	Classical7    string
	Classical8    string
	ClassicalLong string
	ChordChanges1 string
	ChordChanges2 string
	ChordChanges3 string
	ChordChanges4 string
	ChordChanges5 string
	ChordChanges6 string
	ChordChanges7 string
	ChordChanges8 string
}

func generateLead() {
	riffs := getTitles("lead/riffs.json")
	classical := getTitles("lead/classical.json")
	classicalLong := getTitles("lead/classicalLong.json")
	chordChanges := getTitles("lead/chordChanges.json")
	sds := make([]leadSession, 7)
	for i := range sds {
		sds[i].Riff1 = riffs.next()
		println(sds[i].Riff1)
		sds[i].Riff2 = riffs.next()
		sds[i].Riff3 = riffs.next()
		sds[i].Riff4 = riffs.next()
		sds[i].Classical1 = classical.next()
		sds[i].Classical2 = classical.next()
		sds[i].Classical3 = classical.next()
		sds[i].Classical4 = classical.next()
		sds[i].Classical5 = classical.next()
		sds[i].Classical6 = classical.next()
		sds[i].Classical7 = classical.next()
		sds[i].Classical8 = classical.next()
		sds[i].ClassicalLong = classicalLong.next()
		sds[i].ChordChanges1 = chordChanges.next()
		sds[i].ChordChanges2 = chordChanges.next()
		sds[i].ChordChanges3 = chordChanges.next()
		sds[i].ChordChanges4 = chordChanges.next()
		sds[i].ChordChanges5 = chordChanges.next()
		sds[i].ChordChanges6 = chordChanges.next()
		sds[i].ChordChanges7 = chordChanges.next()
		sds[i].ChordChanges8 = chordChanges.next()
	}
	for i := range sds {
		applyTemplate(sds[i], "lead", time.Weekday(i))
	}
}
