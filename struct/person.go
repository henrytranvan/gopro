package main

import "fmt"

type person struct {
	name string
}

func (p *person) Talk() {
	fmt.Println("My name is ", p.name)
	p.name = "Lisa"
}

func main() {
	alex := person{name: "Alex"}

	alex.Talk()
	alex.Talk()
}
