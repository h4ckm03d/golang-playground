package main

import (
	"sync"
	"testing"
)

type Person struct {
	Name  string
	Age   int
	Title string
}

var pool = sync.Pool{
	New: func() interface{} {
		return &Person{}
	},
}

func BenchmarkObjectCreationNoPool(b *testing.B) {
	var person *Person

	for n := 0; n < b.N; n++ {
		person = &Person{
			Title: "The Art of Computer Programming, Vol. 1",
			Name:  "Donald E. Knuth",
			Age:   72,
		}
	}

	_ = person
}

func BenchmarkObjectCreationPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		person := pool.Get().(*Person)
		person.Title = "The Art of Computer Programming, Vol. 1"
		person.Name = "Donald E. Knuth"
		person.Age = 72

		pool.Put(person)
	}
}
