// Single Responsability Principle
/*
	One responsability
	One reason to change
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ....
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// Everythin is okay so far -------
// Separation of Concerns ----

// GoD Object - Anti Pattern
// When you take everythin in put into a single package/class

// Breaking the Single Responsability Principle, -----
/*
we are adding to the Journal concerns to save, load data using io operations
the responsability of the Journal is to manage the entris not deal with IO operations/persistences
persistences could use another package to handle that because we could have others Entities in our system
tha also would need to deal with persistence,
so we can create a package responsible to that and share among our system needss
*/
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	// ...
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	// ...
}

func main() {
	journal := Journal{}
	persistence := Persistence{}

	journal.AddEntry("I learned something new today")
	journal.AddEntry("I teached someone today")

	persistence.SaveToFile("my_journal.txt", journal.String())
}

// Persistence Responsability

type Persistence struct {
	lineSeperator string
}

func (p *Persistence) SaveToFile(filename string, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0644)
}
