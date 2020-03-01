package engine

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// Server App server
type Server struct {
	pattern   string
	messages  []*Message
	clients   map[int]*Client
	addCh     chan *Client
	delCh     chan *Client
	sendAllCh chan *Message
	doneCh    chan bool
	errCh     chan error
}

// NewServer  Create new app server
func NewServer(pattern string) *Server {
	messages := []*Message{}
	clients := make(map[int]*Client)
	addCh := make(chan *Client)
	delCh := make(chan *Client)
	sendAllCh := make(chan *Message)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &Server{
		pattern,
		messages,
		clients,
		addCh,
		delCh,
		sendAllCh,
		doneCh,
		errCh,
	}
}

// Add new player
func (s *Server) Add(c *Client) {
	s.addCh <- c
}

// Del player go out
func (s *Server) Del(c *Client) {
	s.delCh <- c
}

// SendAll – message to channel
func (s *Server) SendAll(msg *Message) {
	s.sendAllCh <- msg
}

// Done with servers work
func (s *Server) Done() {
	s.doneCh <- true
}

// Err – an error happened
func (s *Server) Err(err error) {
	s.errCh <- err
}

// Broadcast send messages
func (s *Server) sendAll(msg *Message) {
	for _, c := range s.clients {
		c.Write(msg)
	}
}

// Shows message history to user
func (s *Server) sendPastMessagees(c *Client) {
	for _, msg := range s.messages {
		c.Write(msg)
	}
}

// Listen and serve.
// It serves client connection and broadcast request.
func (s *Server) Listen() {
	log.Println("Listening server...")

	// websocket handler
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()
		client := NewClient(ws, s)
		s.Add(client)
		client.Listen()
	}

	http.Handle(s.pattern, websocket.Handler(onConnected))
	log.Println("Created handler")

	for {
		select {
		// Add new player
		case c := <-s.addCh:
			log.Println("Added new player")
			s.clients[c.id] = c
			log.Println("Now", len(s.clients), "players connected.")
			s.sendPastMessagees(c)
		// Delete a player
		case c := <-s.delCh:
			log.Println("Delete a player")
			delete(s.clients, c.id)
		// Broadcast send messages for all clients
		case msg := <-s.sendAllCh:
			log.Println("Send to all:", msg)
			s.messages = append(s.messages, msg)
			s.sendAll(msg)
		// Catched an error
		case err := <-s.errCh:
			log.Println("Error:", err.Error())
		// Shut down server
		case <-s.doneCh:
			return
		}
	}
}
