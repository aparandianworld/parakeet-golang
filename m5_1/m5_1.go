package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("package: m5_1")
	fmt.Println("version: 0.0.1")

	printCurrentTime()

	result := simpleAddition(1, 2)
	fmt.Println(result)

	total := improvedAddition(1, 2)
	fmt.Println(total)

	sum := completeAddition(1, 2, 3, 4, 5)
	fmt.Println(sum)

}

func printCurrentTime() {
	now := time.Now()
	fmt.Println(now)
}

func simpleAddition(num1 int, num2 int) int {
	return num1 + num2
}

func improvedAddition(num1, num2 int) int {
	return num1 + num2
}

func completeAddition(numbers ...int) int {
	total := 0

	for _, num := range numbers {
		total += num
	}

	return total
}
