package engine

// Card is a card of deck
type Card struct {
	value float32
	name string
}

// NewCard is a constructor
func NewCard(value float32, name string) *Card {
	return &Card{
		value,
		name,
	}
}
