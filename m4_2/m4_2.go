package main

import (
	"fmt"
)

func main() {
	fmt.Println("package: m4_2")
	fmt.Println("version: 0.0.1")

	rgb := []string{"Red", "Green", "Blue"}

	// old
	for i := 0; i < len(rgb); i++ {
		fmt.Println(rgb[i])
	}

	// modern
	colors := []string{"Black", "White"}
	for _, color := range colors {
		fmt.Println(color)
	}

	// map
	states := map[string]string{
		"AL": "Alabama",
		"AK": "Alaska",
		"AZ": "Arizona",
	}

	for key, value := range states {
		fmt.Println("Key: ", key, ", Value: ", value)
	}

	value := 0
	sum := 0

	// no while loop
	for value < 10 {
		sum += value
		fmt.Println("Value: ", value, ", Sum: ", sum)
		value++
	}

	// goto (legacy)
	target := 24

	for target < 1000 {
		target += target
		if target > 219 {
			goto EOF
		}
	}
	EOF: fmt.Println("Target reached: ", target)
}
