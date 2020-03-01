package engine

// Message model
type Message struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

// String – model to string
func (m *Message) String() string {
	return m.Author + " says: " + m.Body
}
