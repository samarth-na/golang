package basics

import (
	"fmt"
	"time"
)

type Employee struct {
	firstName string
	lastName  string
	age       int
}

type country struct {
	name string

}
type book struct {
	title    string
	auther   string
	numpages int
	issaved  bool
	savedat  time.Time
}

func saved(book book) {
	book.issaved = true
}

func (b *book) titlee() {
	b.title = "name"
}

func notmain() {
	var dune book
	dune.title = "dune"
	ethan := Employee{
		"Sam",
		"Smith",
		25,
	}
	fmt.Println(ethan)
}
