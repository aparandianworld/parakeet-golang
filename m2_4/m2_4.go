package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("package: m2_4")
	fmt.Println("version: 0.0.1")


	f1, f2, f3 := 23.5, 65.1, 76.3
	sumFloat := f1 + f2 + f3
	fmt.Println("f1: ", f1)
	fmt.Println("f2: ", f2)
	fmt.Println("f3: ", f3)
	fmt.Println("sumFloat: ", sumFloat) // sumFloat:  164.89999999999998

	sumFloat = math.Round(sumFloat * 100) / 100
	fmt.Println("sumFloat rounded: ", sumFloat) // sumFloat rounded:  164.9

	pi := math.Pi
	fmt.Println("pi: ", pi)

	radius := 5.2
	area := pi * math.Pow(radius, 2)
	fmt.Printf("area: %.2f", area)
}
