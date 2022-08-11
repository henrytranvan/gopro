package main

import "fmt"

func main() {
	fmt.Println("Enter player name")
	var player1, player2 string
	fmt.Scan(&player1, &player2)
	fmt.Println(player1, player2)
}
