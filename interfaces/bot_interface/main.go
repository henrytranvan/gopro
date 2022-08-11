package main

import "fmt"

type bot interface {
	getGreetingBot() string
}

type englishBot struct{}
type spanishBot struct{}

func printMessage(b bot) {
	fmt.Println(b.getGreetingBot())
}

func (englishBot) getGreetingBot() string {
	return "Hello"
}

func (spanishBot) getGreetingBot() string {
	return "Hallo"
}

func main() {
	en := englishBot{}
	sp := spanishBot{}
	printMessage(en)
	printMessage(sp)
}
