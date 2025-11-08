package main

import (
	"fmt"
)

func main() {
	fmt.Println("package: m3_2")
	fmt.Println("version: 0.0.1")

	var num int = 42
	var p *int = &num

	fmt.Println("p: ", *p)
	fmt.Println("&p: ", &p)
	fmt.Println("num: ", num)

	f := 24.2
	fp := &f

	fmt.Println("f: ", f)
	fmt.Println("fp: ", *fp)
	fmt.Println("&fp: ", &fp)

}
