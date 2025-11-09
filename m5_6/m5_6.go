package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	fmt.Println("package: m5_6")
	fmt.Println("version: 0.0.1")

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)
	req.Header.Set("User-Agent", "Safari/537.36") // if unset, server returns HTTP 403 forbidden error
	resp, err := client.Do(req)
	checkError(err)

	defer resp.Body.Close()

	fmt.Printf("response body type: %T\n", resp.Body)

	body, err := io.ReadAll(resp.Body)
	checkError(err)

	fmt.Println("response body: \n", string(body))
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("network operation failed: %v", err)
	}
}
