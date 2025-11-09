package main

import (
	"fmt"
)

func main() {
	fmt.Println("package: m5_2")
	fmt.Println("version: 0.0.1")

	aaron := Student{Name: "Aaron", Age: 35, GPA: 3.82, Major: "Computer Science"}
	aaron.introduceYourself()

	aaron.changeMajor("Mathematics")
	aaron.introduceYourself()

}

type Student struct {
	Name  string
	Age   int
	GPA   float64
	Major string
}

func (s Student) introduceYourself() {
	fmt.Println("Hello, my name is", s.Name, ".I am ", s.Age, "years old, my major is", s.Major, "and my GPA is", s.GPA)
}

func (s *Student) changeMajor(newMajor string) { // pointer receiver to change the value of fields in the struct
	s.Major = newMajor
}
