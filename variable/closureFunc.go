package main

import "fmt"

func main() {
	add := func(x, y int) int {
		fmt.Println("This is a closure function in Go")
		return x + y
	}

	fmt.Println("Closure add function result:", add(4, 5))
}
