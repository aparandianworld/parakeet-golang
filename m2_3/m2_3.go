package main

import (
	"fmt"
)

func main() {
	fmt.Println("package: m2_3")
	fmt.Println("version: 0.0.1")

	i1, i2, i3 := 10, 20, 30
	sumInt := i1 + i2 + i3

	fmt.Println("i1: ", i1)
	fmt.Println("i2: ", i2)
	fmt.Println("i3: ", i3)
	fmt.Println("sumInt: ", sumInt)

	f1, f2, f3 := 1.1, 2.2, 3.3
	sumFloat := f1 + f2 + f3
	fmt.Println("f1: ", f1)
	fmt.Println("f2: ", f2)
	fmt.Println("f3: ", f3)
	fmt.Println("sumFloat: ", sumFloat)

	product := float64(i1) * f3 // avoid type mismatch
	fmt.Println("product: ", product)
	
}
