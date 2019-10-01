package main

import (
	"io/ioutil"
	"testing"
	"text/template"
)

var bookTemplate string = `
Title: {{.Title}},
Author: {{.Author}}
{{ if .Pages}}
Number of pages: {{ .Pages }}.
{{ end }}
{{ range .Chapters }}
{{ . }}, 
{{ end }}
`

var book *Book = &Book{
	Title:    "The Art of Computer Programming, Vol. 3",
	Author:   "Donald E. Knuth",
	Pages:    800,
	Chapters: []string{"Sorting", "Searching"},
}

func BenchmarkExecute(b *testing.B) {
	t := template.Must(template.New("book").Parse(bookTemplate))

	for n := 0; n < b.N; n++ {
		err := t.Execute(ioutil.Discard, book)
		if err != nil {
			panic(err)
		}
	}
}
