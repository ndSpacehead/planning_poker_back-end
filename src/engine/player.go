package engine

// Player is a user who envolved in estimation process
type Player struct {
	name string
	deck *Deck
}

// NewPlayer is a constructor
func NewPlayer(name string) *Player{
	return &Player{
		name,
		NewDeck(),
	}
}
