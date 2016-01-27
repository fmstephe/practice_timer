package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
	"text/template/parse"
	"time"
)

func main() {
	applyTemplate("rhythm")
	applyTemplate("lead")
}

func applyTemplate(category string) string {
	bytes, err := ioutil.ReadFile(category + "/template.json")
	if err != nil {
		log.Fatal(err)
	}
	t := template.Must(template.New("session").Parse(string(bytes)))
	tMan := newTitleManager(t, category)
	for d := time.Sunday; d <= time.Saturday; d++ {
		tMan.makeValues()
		err = os.MkdirAll("../practice/"+category+"/", 0777)
		if err != nil {
			log.Fatal(err)
		}
		fd, err := os.Create("../practice/" + category + "/" + d.String() + ".json")
		if err != nil {
			log.Fatal(err)
		}
		err = t.Execute(fd, tMan.varMap)
		if err != nil {
			log.Fatal(err)
		}
	}
	return ""
}

type titleManager struct {
	t        *template.Template
	category string
	titleMap map[string]*titles
	varMap   map[string]string
}

func newTitleManager(t *template.Template, category string) *titleManager {
	return &titleManager{
		t:        t,
		category: category,
		titleMap: make(map[string]*titles),
		varMap:   make(map[string]string),
	}
}

func (m *titleManager) makeValues() {
	m.varMap = make(map[string]string)
	for _, n := range m.t.Root.Nodes {
		if n.Type() == parse.NodeAction {
			m.makeValue(n.String(), m.category)
		}
	}
}

func (m *titleManager) makeValue(templateVar, category string) {
	varName := strings.Trim(templateVar, "{}.")
	if _, ok := m.varMap[varName]; ok {
		return
	}
	titleName := strings.TrimRight(varName, "0123456789")
	if titles, ok := m.titleMap[titleName]; ok {
		m.varMap[varName] = titles.next()
		m.makeValue(templateVar, category)
		return
	}
	titles := getTitles(category + "/" + titleName + ".txt")
	m.titleMap[titleName] = titles
	m.makeValue(templateVar, category)
}

type titles struct {
	idx    int
	Titles []string
}

func (t *titles) next() string {
	title := t.Titles[t.idx%len(t.Titles)]
	t.idx++
	return title
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
