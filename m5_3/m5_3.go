package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("package: m5_3")
	fmt.Println("version: 0.0.1")

	go mockFileIO() // This is a goroutine (non-blocking main thread)
	fmt.Println("main function 1")

	time.Sleep(5 * time.Second) // give goroutine time to finish

	fmt.Println("main function 2")

	go func(message string) {
		time.Sleep(2 * time.Second) // give goroutine time to finish
		fmt.Println("message: ", message)
	}("message from anonymous goroutine.")

	fmt.Println("main function 3")

	time.Sleep(5 * time.Second) // give goroutine time to finish
}

func mockFileIO() {
	time.Sleep(2 * time.Second)
	fmt.Println("mockFileIO function")
}
