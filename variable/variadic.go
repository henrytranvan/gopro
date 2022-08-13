package main

import "fmt"

func addMulti(args ...int) int {
	sum := 0
	for _, val := range args {
		sum += val
	}
	return sum
}

func main() {
	fmt.Println("Sum: ", addMulti(4, 5, 6, 3, 2, 2, 4))
}
