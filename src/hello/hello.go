package main

import "fmt"

func main() {
	var hello = "hello\n"
	var world = "world\n"

	defer fmt.Printf(world)
	fmt.Printf(hello)

	a := 10
	fmt.Println(a)
}
