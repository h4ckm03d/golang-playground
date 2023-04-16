package main

import (
	"testing"
)

type Student struct {
	Name string
	Age  int
}

type Getter interface {
	GetName() string
	GetAge() int
}

func (p Student) GetName() string {
	return p.Name
}

func (p Student) GetAge() int {
	return p.Age
}

func getStudent() Getter {
	return Student{Name: "Bob", Age: 30}
}

func returnObject() Student {
	return Student{Name: "Bob", Age: 30}
}

func returnPinterObject() *Student {
	return &Student{Name: "Bob", Age: 30}
}

func returnInterface() Getter {
	return Student{Name: "Bob", Age: 30}
}

func BenchmarkReturnObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		returnObject()
	}
}

func BenchmarkReturnPointerObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		returnPinterObject()
	}
}

func BenchmarkReturnInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		returnInterface()
	}
}
