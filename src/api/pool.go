package api

import "fmt"

// Pool is struct to hold details of connection pool for websocket
type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Games      map[string][]*Client
	Broadcast  chan Envelope
}

// NewPool instantiates and returns a new Pool obj
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Games:      make(map[string][]*Client),
		Broadcast:  make(chan Envelope),
	}
}

// Start is method of Pool which will take an action based on the client channel
func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Games[client.id] = append(pool.Games[client.id], client)
			fmt.Printf("Total number of clients in pool: %d\n", len(pool.Games))
			fmt.Printf("Number of clients in Game %s: %d\n", client.id, len(pool.Games[client.id]))

			for _, client := range pool.Games[client.id] {
				fmt.Printf("%v Client Joined Game %s\n", client, client.id)
			}
			break
		case client := <-pool.Unregister:
			for i, c := range pool.Games[client.id] {
				if c == client {
					fmt.Printf("Deleting client: %v\n", c)
					pool.Games[client.id] = append(pool.Games[client.id][:i], pool.Games[client.id][i+1:]...)
				}
				fmt.Printf("%v Client Left Game %s\n", client, client.id)
			}
			break
		case event := <-pool.Broadcast:
			fmt.Printf("Sending message to %d clients in Game %s\n", len(pool.Games[event.ClientID]), event.ClientID)

			for i, client := range pool.Games[event.ClientID] {
				fmt.Printf("Message %d sent\n", i+1)
				if err := client.socket.WriteJSON(event.Message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
