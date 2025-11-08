package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("package: m4_1")
	fmt.Println("version: 0.0.1")

	num := 42
	var answer string

	// conditional similar to c language (no need for parentheses)
	if num < 0 {
		answer = "Less than 0"
	} else if num > 0 {
		answer = "Greater than 0"
	} else {
		answer = "Equal to 0"
	}

	fmt.Println(answer)

	today := time.Now().Weekday()
	fmt.Println("today is: ", today)

	dayNumber := int(today)
	fmt.Println("day number is: ", dayNumber)

	// switch no need for break
	switch dayNumber {
	case 0:
		fmt.Println("Sunday")
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
		fallthrough
	default:
		fmt.Println("Default code executed here...")
	}

}
