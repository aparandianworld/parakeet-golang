package main

import (
	"fmt"
)

// similar to c structs (group related values together)
type Student struct {
	Name   string
	Age    int
	Gender string
	School string
	Major  string
}

func main() {
	fmt.Println("package: m3_6")
	fmt.Println("version: 0.0.1")

	aaron := Student{"Aaron", 22, "Male", "Stanford University", "Software Engineering"}
	fmt.Println(aaron)
	fmt.Printf("%+v\n", aaron) // include field names
	fmt.Printf("Name: %s\n", aaron.Name)
	fmt.Printf("Age: %d\n", aaron.Age)
	fmt.Printf("Gender: %s\n", aaron.Gender)
	fmt.Printf("School: %s\n", aaron.School)
	fmt.Printf("Major: %s\n", aaron.Major)
}
