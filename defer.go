package main

import "fmt"

func main() {
	fmt.Println("Hello")
	panic("Panic Attack")
	fmt.Println("World")
	defer fmt.Println("Execute Defer")
}
