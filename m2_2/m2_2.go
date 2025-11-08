package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello", name)

	fmt.Println("Enter your age: ")
	age, _ := reader.ReadString('\n')
	fmt.Println("Your age: ", age)

	fmt.Println("Enter a float number: ")
	num, _ := reader.ReadString('\n')
	number, err := strconv.ParseFloat(strings.TrimSpace(num), 64)
	if err != nil {
		fmt.Println("Invalid number")
		return
	} else {
		fmt.Println("Your number: ", number)
	}

}
