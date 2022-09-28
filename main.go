package main

func main() {
	cards := deck{newCard(), "Six of heart"}
	cards = append(cards, "Seven of heart")
	cards.print()
}

func newCard() string {
	return "Five of heart"
}
