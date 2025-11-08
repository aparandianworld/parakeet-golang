package main

import (
	"fmt"
)

func main() {
	fmt.Println("package: m3_3")
	fmt.Println("version: 0.0.1")

	var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"

	fmt.Println(colors)
	fmt.Println("length: ", len(colors))
	fmt.Println("capacity: ", cap(colors))
	fmt.Println("colors[0]: ", colors[0])

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println("length: ", len(numbers))
	fmt.Println("capacity: ", cap(numbers))
	fmt.Println("numbers[0]: ", numbers[0])

	// invalid append operation on fixed-size array
	// numbers = append(numbers, 6) 
	// fmt.Println(numbers)
}
