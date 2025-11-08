package main

import (
	"fmt"
)

func main() {
	fmt.Println("package: m3_1")
	fmt.Println("version: 0.0.1")

	n := new(int) // new allocates memory but does not initialize beyond zero values and returns a pointer
	fmt.Println(n) 

	m := make(map[string]int) // make allocates and initializes memory
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println("langth: ", len(m))
	fmt.Println(m)

	fmt.Println(m["one"])
	fmt.Println(m["two"])
	fmt.Println(m["three"])
	
}