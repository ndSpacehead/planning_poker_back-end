package engine

// Deck is a collection of cards
type Deck struct {
	cards []*Card
}

// NewDeck is a constructor
func NewDeck() *Deck {
	cards := []*Card{
		NewCard(0, "0"),
		NewCard(0.5, "1/2"),
		NewCard(1, "1"),
		NewCard(2, "2"),
		NewCard(3, "3"),
		NewCard(5, "5"),
		NewCard(8, "8"),
		NewCard(13, "13"),
		NewCard(20, "20"),
		NewCard(40, "40"),
		NewCard(100, "100"),
		NewCard(1000, "?"),
		NewCard(-1, "coffee"),
	}
	return &Deck{
		cards,
	}
}
