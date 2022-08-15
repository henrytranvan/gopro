package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	data []interface{}
	top  int
}

func (s *stack) push(element interface{}) {
	s.top++
	s.data = append(s.data, element)
}

func (s *stack) pop() interface{} {
	if len(s.data) > 0 {
		last := s.data[s.top-1]
		s.top--
		s.data = s.data[0:s.top]
		return last
	}
	return nil
}

func (s stack) print() {
	for i, el := range s.data {
		fmt.Println(i, el)
	}
}

func (s stack) length() int {
	return s.top
}

func (s *stack) clear() {
	s.top = 0
	s.data = []interface{}{}
}

func (s stack) peek() interface{} {
	if len(s.data) > 0 {
		return s.data[s.top-1]
	}
	return nil
}

func main() {
	s := stack{}
	s.push(23)
	s.push(55)
	s.push(45)
	s.push(65)
	s.print()
	numbers := s.pop()
	fmt.Println(numbers)
	s.print()

	number := 19

	fmt.Printf("Number %d converted to base %d is %s\n", number, 2, numberTransformBase(number, 2))
	fmt.Printf("Number %d converted to base %d is %s\n", number, 8, numberTransformBase(number, 8))

}

func numberTransformBase(number int, base int) string {
	s := stack{}
	for {
		s.push(number % base)
		number = number / base

		if number <= 0 {
			break
		}
	}
	converted := ""

	for s.length() > 0 {
		converted = converted + strconv.Itoa(s.pop().(int))
	}

	return converted
}
