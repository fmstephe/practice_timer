package main

import "time"

type rhythmSession struct {
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
	NeckRhythm1 string
	NeckRhythm2 string
	SingleNote4 string
	SingleNote5 string
	//
	Barre1      string
	NeckRhythm3 string
	SingleNote6 string
	SingleNote7 string
}

func generateRhythm() {
	barre := getTitles("rhythm/barre.txt")
	oneChordRhythm := getTitles("rhythm/oneChordRhythm.txt")
	twoChordRhythm := getTitles("rhythm/twoChordRhythm.txt")
	neckRhythm := getTitles("rhythm/neckRhythm.txt")
	singleNote := getTitles("rhythm/singleNote.txt")
	sds := make([]rhythmSession, 7)
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
		sds[i].NeckRhythm1 = neckRhythm.next()
		sds[i].NeckRhythm2 = neckRhythm.next()
		sds[i].SingleNote4 = singleNote.next()
		sds[i].SingleNote5 = singleNote.next()
		//
		sds[i].Barre1 = barre.next()
		sds[i].NeckRhythm3 = neckRhythm.next()
		sds[i].SingleNote6 = singleNote.next()
		sds[i].SingleNote7 = singleNote.next()
	}
	for i := range sds {
		applyTemplate(sds[i], "rhythm", time.Weekday(i))
	}
}
