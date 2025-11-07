package main

import (
	"fmt"
)

func main() {

	// implicit values

	str1 := "The quick brown fox"
	str2 := "jumps over the lazy dog"

	age := 3

	fmt.Println(str1, str2)
	str_length, err := fmt.Println("Age: ", age)

	if err == nil {
		fmt.Println("String length: ", str_length)
		fmt.Println("Error: ", err)
		fmt.Printf("Value of Age: %v\n", age)
		fmt.Printf("Type of String Length: %T\n", str_length)
		fmt.Printf("Type of Error: %T\n", err)
		fmt.Printf("Type of Age: %T\n", age)
	}

}
