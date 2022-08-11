package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(first int, second int) int {
	return first + second
}

func main() {
	first, _ := strconv.Atoi(os.Args[1])
	second, _ := strconv.Atoi(os.Args[2])
	sum := add(first, second)
	fmt.Println(sum)
}
