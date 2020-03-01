package engine

// Room is a place where players estimate
type Room struct {
	name    string
	players []*Player
}

// NewRoom is a constructor
func NewRoom(name string) *Room {
	return &Room{
		name,
		[]*Player{},
	}
}
