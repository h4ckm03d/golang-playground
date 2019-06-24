package main

import (
	"testing"
)

func BenchmarkA(b *testing.B) {
	p := Person{Name: "bajindul", Age: 32}
	for n := 0; n < b.N; n++ {
		NewStudentA(p)
	}
}

func BenchmarkB(b *testing.B) {
	p := Person{Name: "bajindul", Age: 32}
	for n := 0; n < b.N; n++ {
		NewStudentB(&p)
	}
}

func BenchmarkC(b *testing.B) {
	p := Person{Name: "bajindul", Age: 32}
	for n := 0; n < b.N; n++ {
		NewStudentC(&p)
	}
}

func BenchmarkAParalel(b *testing.B) {
	p := Person{Name: "bajindul", Age: 32}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			NewStudentA(p)
		}
	})
}

func BenchmarkBParalel(b *testing.B) {
	p := Person{Name: "bajindul", Age: 32}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			NewStudentB(&p)
		}
	})
}

func BenchmarkCParalel(b *testing.B) {
	p := Person{Name: "bajindul", Age: 32}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			NewStudentC(&p)
		}
	})
}
