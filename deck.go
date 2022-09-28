package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardsSuits := deck{"Chuon", "Bich", "Ro", "Co"}
	cardsValues := deck{"Mot", "Hai", "Ba", "Bon"}

	for _, value := range cardsValues {
		for _, suit := range cardsSuits {
			cards = append(cards, value+" "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
