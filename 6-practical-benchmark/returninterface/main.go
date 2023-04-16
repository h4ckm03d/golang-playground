package main

import (
	"fmt"
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

func main() {
	s := getStudent()
	fmt.Println(s.GetAge(), s.GetName())
}
