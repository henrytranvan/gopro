package main

import "fmt"

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

func main() {
	s := stack{}
	s.push(23)
	s.push(55)
	s.push(45)
	s.push(65)
	s.print()
	number := s.pop()
	fmt.Println(number)
	s.print()
}
