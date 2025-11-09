package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("package: m5_5")
	fmt.Println("version: 0.0.1")

	fileName := "./example.txt"
	file, err := os.Create(fileName)
	checkError(err)
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("warning - failed to close the file: %v", err)
		}
	}()

	length, err := io.WriteString(file, "This is an example of writing to a file in Go\n")
	checkError(err)
	fmt.Printf("length of data written: %d\n", length)

	readFile(fileName)
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("file operation failed: %v", err)
	}
}

func readFile(fileName string) {
	file, err := os.Open(fileName)
	checkError(err)
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("warning - failed to close the file: %v", err)
		}
	}()
	data, err := io.ReadAll(file)
	checkError(err)

	fmt.Println("file opened successfully for reading...")
	fmt.Printf("file content: %v\n", string(data))
}
