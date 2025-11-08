package main

import "fmt"

type Color struct {
	Name string
	Hex  int
}

func slicesToObjects(colorNames []string, hexValues []int) []Color {
	colors := make([]Color, len(colorNames))

	for i := range colorNames {
		colors[i] = Color{
			Name: colorNames[i],
			Hex:  hexValues[i],
		}
	}

	return colors
}

func main() {
	fmt.Println("package: m3_7")
	fmt.Println("version: 0.0.1")

	colorNames := []string{"Red", "Green", "Blue"}
	hexValues := []int{0xFF0000, 0x00FF00, 0x0000FF}
	colors := slicesToObjects(colorNames, hexValues)
	fmt.Println(colors)
}
