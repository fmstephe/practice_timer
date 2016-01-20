package tab

import (
	"fmt"
	"strconv"
	"strings"
)

// "1:2|2:3|3:0"
// means string 1 fret 2, string 2 fret 3, string 3 open
// - means an empty column
type Column struct {
	Notes map[string]string
}

func (c *Column) Print(stringNum string) string {
	note, ok := c.Notes[stringNum]
	if !ok {
		return "-"
	}
	return note
}

func ParseColumn(cStr string) Column {
	c := Column{
		Notes: make(map[string]string),
	}
	if cStr != "-" {
		notes := strings.Split(cStr, "|")
		for i := range notes {
			pair := strings.Split(notes[i], ":")
			c.Notes[pair[0]] = pair[1]
		}
	}
	return c
}

// "[...,...,...]"
// Contains a bar with three columns, separated by ','
type Bar struct {
	Columns []Column
}

func (b *Bar) Print(stringNum string) string {
	notes := ""
	for _, c := range b.Columns {
		notes += c.Print(stringNum)
	}
	return notes + "|"
}

func ParseBar(barStr string) Bar {
	innerStr := barStr[1 : len(barStr)-1]
	columnStrs := strings.Split(innerStr, ",")
	b := Bar{}
	for _, columnStr := range columnStrs {
		b.Columns = append(b.Columns, ParseColumn(columnStr))
	}
	return b
}

// "[...][...][...]"
// Contains a motif with three bars
type Motif struct {
	Bars []Bar
}

func (m *Motif) Print(stringNum string) string {
	notes := ""
	for _, b := range m.Bars {
		notes += b.Print(stringNum)
	}
	return notes
}

func (m *Motif) String() string {
	str := ""
	for i := 6; i >= 1; i-- {
		str += m.getStringName(i) + m.Print(strconv.Itoa(i)) + "|\n"
	}
	return str
}

func (m *Motif) getStringName(i int) string {
	switch i {
	case 1:
		return "E||"
	case 2:
		return "A||"
	case 3:
		return "D||"
	case 4:
		return "G||"
	case 5:
		return "B||"
	case 6:
		return "e||"
	default:
		panic(fmt.Sprintf("Bad string number %d", i))
	}
}

func ParseMotif(motifStr string) Motif {
	barStrs := strings.SplitAfter(motifStr, "]")
	m := Motif{}
	for _, barStr := range barStrs {
		if barStr != "" {
			m.Bars = append(m.Bars, ParseBar(barStr))
		}
	}
	return m
}
