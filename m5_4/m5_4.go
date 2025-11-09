package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("package: m5_4")
	fmt.Println("version: 0.0.1")

	value1 := "10"
	value2 := "5.5"
	operation := "+"
	result := calculate(value1, value2, operation)
	fmt.Println(result)
}

func calculate(input1 string, input2 string, operation string) float64 {
	input1f64 := convertInputToValue(input1)
	input2f64 := convertInputToValue(input2)

	switch operation {
	case "+":
		return addValues(input1f64, input2f64)
	case "-":
		return subtractValues(input1f64, input2f64)
	case "*":
		return multiplyValues(input1f64, input2f64)
	case "/":
		return divideValues(input1f64, input2f64)
	default:
		return 0
	}
}

func convertInputToValue(input string) float64 {
	value, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	return value
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
