package main

import (
	"fmt"
)

func remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	fmt.Println("package: m3_4")
	fmt.Println("version: 0.0.1")

	colors := []string{"red", "green", "blue"} // This is a slice

	fmt.Println(colors)
	fmt.Println("length: ", len(colors))
	fmt.Println("capacity: ", cap(colors))
	fmt.Println("colors[0]: ", colors[0])
	fmt.Println("colors[1]: ", colors[1])
	fmt.Println("colors[2]: ", colors[2])

	numbers := make([]int, 0, 5) // This is a slice (memory efficient)
	numbers = append(numbers, 1, 2, 3, 4, 5)
	fmt.Println(numbers)
	fmt.Println("length: ", len(numbers))
	fmt.Println("capacity: ", cap(numbers))
	fmt.Println("numbers[0]: ", numbers[0])
	fmt.Println("numbers[1]: ", numbers[1])
	fmt.Println("numbers[2]: ", numbers[2])
	fmt.Println("numbers[3]: ", numbers[3])
	fmt.Println("numbers[4]: ", numbers[4])

	// add colors
	fmt.Println("adding more colors...")
	colors = append(colors, "yellow", "white", "black")
	fmt.Println(colors)
	fmt.Println("length: ", len(colors))
	fmt.Println("capacity: ", cap(colors))
	fmt.Println("colors[0]: ", colors[0])
	fmt.Println("colors[1]: ", colors[1])
	fmt.Println("colors[2]: ", colors[2])
	fmt.Println("colors[3]: ", colors[3])
	fmt.Println("colors[4]: ", colors[4])
	fmt.Println("colors[5]: ", colors[5])

	// remove numbers
	fmt.Println("removing number at index 4...")
	numbers = remove(numbers, 4)
	fmt.Println(numbers)
	fmt.Println("length after remove: ", len(numbers))
	fmt.Println("capacity after remove: ", cap(numbers))
}
