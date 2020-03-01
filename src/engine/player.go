package engine

// Player is a user who envolved in estimation process
type Player struct {
	Client
	name   string
	deck   *Deck
	room   *Room
	doneCh chan bool
}

// NewPlayer is a constructor
func NewPlayer(client Client) *Player {
	doneCh := make(chan bool)

	return &Player{
		client,
		"Player " + string(client.id),
		NewDeck(),
		nil,
		doneCh,
	}
}
