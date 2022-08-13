package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type desk []string

func newDesk() desk {
	cards := desk{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func (d desk) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d desk, handsize int) (desk, desk) {
	return d[:handsize], d[handsize:]
}

func (d desk) toString() string {
	return strings.Join(d, ", ")
}

func (d desk) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 07777)
}

func newDeskFromFile(fileName string) desk {
	strCard, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File dead error: ", err)
		os.Exit(1)
	}
	return strings.Split(string(strCard), ", ")
}

func (d desk) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
