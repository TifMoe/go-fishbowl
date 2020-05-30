package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

// NewRouter will build a new router for the websocket connection
func NewRouter(p *Pool, c GameController) *WebSocketRouter {

	wsr := NewWebSocketRouter(p)
	wsr.Handle("disconnect", Disconnect)

	wsr.Handle("newGame", c.NewGame)
	wsr.Handle("getGame", c.GetGame)
	wsr.Handle("updateGame", c.UpdateGame)
	wsr.Handle("resetGame", c.ResetGame)
	wsr.Handle("startRound", c.StartRound)

	wsr.Handle("newCard", c.NewCard)
	wsr.Handle("usedCard", c.MarkCardUsed)
	wsr.Handle("getRandomCard", c.GetRandomCard)

	return wsr
}

// Handler is a type representing functions which resolve requests.
type Handler func(*Client, interface{})

// Event is a type representing request names.
type Event string

// WebSocketRouter is a message routing object mapping events to function handlers.
type WebSocketRouter struct {
	rules map[Event]Handler // rules maps events to functions.
	pool  *Pool
}

// NewWebSocketRouter returns an initialized Router.
func NewWebSocketRouter(p *Pool) *WebSocketRouter {
	return &WebSocketRouter{
		rules: make(map[Event]Handler),
		pool:  p,
	}
}

// ServeHTTP creates the socket connection and begins the read routine.
func (rt *WebSocketRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// configure upgrader
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// TODO: don't accept all by default
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	// upgrade connection to socket
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("socket server configuration error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	clientGroup := strings.Split(r.URL.String(), "/ws/")[1]
	client := NewClient(clientGroup, rt.pool, socket, rt.FindHandler)
	rt.pool.Register <- client

	// running method for reading from sockets, in main routine
	client.Read()
}

// FindHandler implements a handler finding function for router.
func (rt *WebSocketRouter) FindHandler(event Event) (Handler, bool) {
	handler, found := rt.rules[event]
	return handler, found
}

// Handle is a function to add handlers to the router.
func (rt *WebSocketRouter) Handle(event Event, handler Handler) {
	// store in to router rules
	rt.rules[event] = handler
}

// Disconnect is handler to disconnect client when "disconnect" event is emitted from client
func Disconnect(cl *Client, data interface{}) {
	log.Printf("Force disconnect client: %s\n", cl.id)
	cl.pool.Unregister <- cl
	cl.socket.Close()
}
