package main

/*func main() {
	//var card string = "Ace of Spades"
	//we use := to define a new variable, not revalue one
	//card = "Five of Diamonds "
	//just use the name of the var and u can revalue

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

// when you wanna return a value, write the datatype too
func newCard() string {
	return "Five of Diamonds"
}
*/
//array-fixed length list of things
//slice- an array that can grow and shrink
//slices can only have one data type across elements--such as all int

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	cardsOne := newDeck()
	cardsOne.saveToFile("my_cards")

	//cardsFile := newDeckFromFile("my_cards")
	//cardsFile.print()

	//How to convert to ascii/byte --
	// greeting := "Hi there!"
	//fmt. Println([]byte(greeting))

	cardsShuffled := newDeck()
	cardsShuffled.shuffle()
	cardsShuffled.print()
}

//Structs - Collection of Properties that are related together
