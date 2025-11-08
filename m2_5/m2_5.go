package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("package: m2_5")
	fmt.Println("version: 0.0.1")

	goLaunchDate := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("goLaunchDate: ", goLaunchDate)
	fmt.Printf("goLaunchDate object: %T\n", goLaunchDate)

	currentTime := time.Now()
	fmt.Println("currentTime: ", currentTime)
	fmt.Printf("currentTime object: %T\n", currentTime)

	formattedTime := currentTime.Format(time.ANSIC)
	fmt.Println("formattedTime: ", formattedTime)
	fmt.Printf("formattedTime object: %T\n", formattedTime)

	tomorrow := currentTime.AddDate(0, 0, 1)
	fmt.Println("tomorrow: ", tomorrow)

	customFormat := "Monday 2006-02-01"
	formattedCurrentTime := currentTime.Format(customFormat)
	fmt.Println("formatted current time: ", formattedCurrentTime)
}