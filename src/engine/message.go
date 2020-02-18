package engine

type Message struct {
	body string `json:"body"`
}

func (m *Message) String() string {
	return m.body
}