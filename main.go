package main

func main() {
	cards := newDeck()

	hand, remain := deal(cards, 6)

	hand.print()
	remain.print()
}

func newCard() string {
	return "Five of heart"
}
