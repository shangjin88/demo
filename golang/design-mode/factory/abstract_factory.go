package main

import "fmt"

type Stu interface {
	Greet()
}

type stu struct {
	Name string
	Age  int
}

func (s stu) Greet() {
	fmt.Printf("Hi! My name is %s", s.Name)
}

func NewStu(name string, age int) Stu {
	return stu{
		Name: name,
		Age:  age,
	}
}

func main() {
	s := NewStu("json", 10)
	s.Greet()
}
