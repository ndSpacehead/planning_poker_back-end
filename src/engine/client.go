package engine

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/websocket"
)

const channelBufSize = 100

var maxId int = 0

// Client model
type Client struct {
	id        int
	ws        *websocket.Conn
	server    *Server
	messageCh chan *Message
	doneCh    chan bool
}

// NewClient is a constructor
func NewClient(ws *websocket.Conn, server *Server) *Client {
	if ws == nil {
		panic("ws cannot be nil")
	}
	if server == nil {
		panic("server cannot be nil")
	}

	maxId++
	messageCh := make(chan *Message, channelBufSize)
	doneCh := make(chan bool)

	return &Client{maxId, ws, server, messageCh, doneCh}
}

// Conn returns websocket connection
func (c *Client) Conn() *websocket.Conn {
	return c.ws
}

func (c *Client) Write(msg *Message) {
	select {
	case c.messageCh <- msg:
	default:
		c.server.Del(c)
		err := fmt.Errorf("client %d is disconnected", c.id)
		c.server.Err(err)
	}
}

// Done is terminate session
func (c *Client) Done() {
	c.doneCh <- true
}

// Listen write and read events
func (c *Client) Listen() {
	go c.listenWrite()
	c.listenRead()
}

func (c *Client) listenWrite() {
	log.Println("Listening write to client")
	for {
		select {
		// Send message to the client
		case msg := <-c.messageCh:
			log.Printf("To client %d send: %s", c.id, msg)
			websocket.JSON.Send(c.ws, msg)
		// Recieve donw request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- true // for listenRead method
			return
		}
	}
}

func (c *Client) listenRead() {
	log.Println("Listening read from client")
	for {
		select {
		// Receive done request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- true // for listenWrite method
			return
		// Read data from websocket connetction
		default:
			var msg Message
			err := websocket.JSON.Receive(c.ws, &msg)
			if err == io.EOF {
				c.doneCh <- true
			} else if err != nil {
				c.server.Err(err)
			} else {
				c.server.SendAll(&msg)
			}
		}
	}
}
