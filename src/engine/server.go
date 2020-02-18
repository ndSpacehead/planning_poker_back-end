package engine

import (
	"net/http"

	"golang.org/x/net/websocket"
)

// Server App server
type Server struct {
	pattern  string
	messages []*Message
	doneCh   chan bool
	errCh    chan error
}

// NewServer  Create new app server
func NewServer(pattern string) *Server {
	messages := []*Message{}
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &Server{
		pattern,
		messages,
		doneCh,
		errCh,
	}
}

// Listen and serve.
// It serves client connection and broadcast request.
func (s *Server) Listen() {
	// websocket handler
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()
	}

	http.Handle(s.pattern, websocket.Handler(onConnected))

	for {
		return
	}
}
