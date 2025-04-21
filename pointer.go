package main

import "fmt"

func main() {
	x := 1
	p := &x
	//or
	// var x int = 1
	// var p *int = &x

	fmt.Println(p)
	fmt.Println(*p)
}
