package api

import (
	"log"

	"github.com/gorilla/websocket"
)

// Envelope is a object used to pass data on sockets from specific clients.
type Envelope struct {
	ClientID string  `json:"client"`
	Message  Message `json:"message"`
}

// Message contains event contents
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

// FindHandler is a type that defines handler finding functions.
type FindHandler func(Event) (Handler, bool)

// Client is a type that reads and writes on sockets.
type Client struct {
	id          string
	send        Envelope
	socket      *websocket.Conn
	findHandler FindHandler
	pool        *Pool
}

// NewClient accepts a socket and returns an initialized Client.
func NewClient(id string, p *Pool, socket *websocket.Conn, findHandler FindHandler) *Client {
	return &Client{
		id:          id,
		socket:      socket,
		findHandler: findHandler,
		pool:        p,
	}
}

// Write receives messages from the channel and writes to the socket.
func (c *Client) Write() {
	msg := c.send
	err := c.socket.WriteJSON(msg.Message)
	if err != nil {
		log.Printf("socket write error: %v\n", err)
	}
}

// Read intercepts messages on the socket and assigns them to a handler function.
func (c *Client) Read() {
	var msg Message

	defer func() {
		c.pool.Unregister <- c
		c.socket.Close()
	}()

	for {
		log.Printf("Reading message from client!")
		// read incoming message from socket
		if err := c.socket.ReadJSON(&msg); err != nil {
			log.Printf("socket read error: %v\n", err)
			break
		}
		log.Printf("Message from client!: %+v\n", msg)

		// assign message to a function handler
		if handler, found := c.findHandler(Event(msg.Name)); found {
			handler(c, msg.Data)
		}
	}
	log.Println("exiting read loop")

	// close interrupted socket connection
	c.socket.Close()
}
