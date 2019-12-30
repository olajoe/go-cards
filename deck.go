package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of stings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spade", "Diamond", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
