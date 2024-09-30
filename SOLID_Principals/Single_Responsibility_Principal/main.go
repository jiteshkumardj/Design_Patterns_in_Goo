package main

import "fmt"

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntery(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d : %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}
func (j *Journal) RemoveEntry(index int) {
	// ...
}

//Journal has responsibility of keeping entries and managing those entries

//lets break the SRP By adding the below function
func (j *Journal) Save(filename string) {

}
func main() {
	// it states that a type should have only one primary responsibility
	// As a result it should have one reason to change
	j := Journal{}
	j.AddEntery("Hello there")
	j.AddEntery("I ate a bug today")
	fmt.Println(j)
}
