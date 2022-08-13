package main

import "fmt"

func main() {

	cards := newDesk()

	fmt.Println(cards)
	fmt.Println("================================")
	cards.print()

	a, b := deal(cards, 3)
	fmt.Println("================================")
	a.print()
	b.print()
	fmt.Println("================================")
	fmt.Println(cards.toString())

	cards.saveToFile("cards.txt")

	newCards := newDeskFromFile("cards.txt")
	fmt.Println("================================")
	newCards.print()

	newCards.shuffle()
	fmt.Println("================================")
	newCards.print()
}
