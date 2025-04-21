package main

import (
	"fmt"
)

func main() {
	type fahrenheit int
	type celsius int

	var f fahrenheit = 32
	var c celsius = 0

	fmt.Println(f, c)
}
