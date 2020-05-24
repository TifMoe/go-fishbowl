package api

import (
	"encoding/json"
	"fmt"
	"github.com/tifmoe/go-fishbowl/src/service"
	"log"
)

// NewGameController will instantiate a new game handler
func NewGameController(svc service.GameService) *controller {
	return &controller{
		Svc: svc,
	}
}

// GameController is interface with methods to interact with redis db
type GameController interface {
	NewGame(*Client, interface{})
	GetGame(*Client, interface{})
	UpdateGame(*Client, interface{})
	ResetGame(*Client, interface{})

	NewCard(*Client, interface{})
	GetRandomCard(*Client, interface{})
	MarkCardUsed(*Client, interface{})
	StartRound(*Client, interface{})
}

// Controller holds service for Game Handlers
type controller struct {
	Svc service.GameService
}

// NewGame is controller for generating new game namespace and instantiating in the database
func (c *controller) NewGame(cl *Client, data interface{}) {
	input := service.TeamInput{}

	err := json.Unmarshal([]byte(data.(string)), &input)
	if err != nil {
		log.Printf("error decoding values: %v", err)
		// TODO serve error response
		return
	}
	// Instantiate new game namespace
	nameSpace, err := c.Svc.NewGame(&input)
	if err != nil {
		log.Printf("error generating new namespace: %v", err)
		// TODO return error
		return
	}

	cl.send = Envelope{
		Message: Message{Name: "newGame", Data: nameSpace},
	}
	cl.Write()
}

// UpdateGame is controller for updating a game with new data
func (c *controller) UpdateGame(cl *Client, data interface{}) {
	input := service.GameInput{}

	err := json.Unmarshal([]byte(data.(string)), &input)
	if err != nil {
		log.Printf("error decoding values: %v", err)
		// TODO serve error response
		return
	}

	game, err := c.Svc.UpdateGame(input.ID, &input)
	if err != nil {
		log.Printf("error updating game %s: %v", input.ID, err)
		// TODO serve error response
		return
	}

	gameData, err := json.Marshal(internalToExternal(game))
	if err != nil {
		log.Printf("error marshalling game struct to json %v: %v", game, err)
		// TODO serve error message
		return
	}

	message := Envelope{
		ClientID: input.ID,
		Message:  Message{Name: "gameState", Data: string(gameData)},
	}
	cl.pool.Broadcast <- message
}

// NewCard is controller for saving a new card to the existing game
func (c *controller) NewCard(cl *Client, data interface{}) {
	card := service.CardInput{}

	err := json.Unmarshal([]byte(data.(string)), &card)
	if err != nil || card.IsEmpty() {
		log.Printf("error decoding values: %v", err)
		// TODO return error
		return
	}

	cardCount, err := c.Svc.SaveCard(card.GameID, &card)
	if err != nil {
		log.Printf("error saving values: %v", err)
		// TODO discern validation and not-found errors, return in message
		return
	}
	message := Envelope{
		ClientID: card.GameID,
		Message:  Message{Name: "cardCount", Data: cardCount},
	}
	cl.pool.Broadcast <- message
}

// GetGame is controller for fetching a specific game
func (c *controller) GetGame(cl *Client, data interface{}) {
	game := GameInput{}

	err := json.Unmarshal([]byte(data.(string)), &game)
	if err != nil || game.ID == "" {
		log.Printf("error decoding values: %v", err)
		// TODO serve error message
		return
	}

	gameInternal, err := c.Svc.GetGame(game.ID)
	if err != nil {
		log.Printf("error fetching cards: %v", err)
		// TODO serve error message
		return
	}

	gameData, err := json.Marshal(internalToExternal(gameInternal))
	if err != nil {
		log.Printf("error marshalling game struct to json %v: %v", gameInternal, err)
		// TODO serve error message
		return
	}
	message := Envelope{
		ClientID: game.ID,
		Message:  Message{Name: "gameState", Data: string(gameData)},
	}
	cl.pool.Broadcast <- message
}

// GetRandomCard is controller to return a random un-used card for a specific game
func (c *controller) GetRandomCard(cl *Client, data interface{}) {
	game := GameInput{}

	err := json.Unmarshal([]byte(data.(string)), &game)
	if err != nil || game.ID == "" {
		log.Printf("error decoding values: %v", err)
		// TODO serve error message
		return
	}

	cards := []Card{}

	internalCard, err := c.Svc.GetRandomCard(game.ID)
	if err != nil {
		log.Printf("error fetching cards: %v", err)
		// TODO return error message
		return
	}
	if internalCard != nil {
		cards = []Card{internalToExternalCard(internalCard)}
	}

	gameInternal, err := c.Svc.GetGame(game.ID)
	if err != nil {
		log.Printf("error fetching cards: %v", err)
		// TODO return error message
		return
	}

	gameExternal := internalToExternal(gameInternal)
	gameExternal.Cards = cards

	gameData, err := json.Marshal(gameExternal)
	if err != nil {
		log.Printf("error marshalling game struct to json %v: %v", gameExternal, err)
		// TODO serve error message
		return
	}

	cl.send = Envelope{
		Message: Message{Name: "randomCard", Data: string(gameData)},
	}
	cl.Write()
}

// MarkCardUsed is controller to update values of a specific card
func (c *controller) MarkCardUsed(cl *Client, data interface{}) {
	input := CardUsedInput{}

	err := json.Unmarshal([]byte(data.(string)), &input)
	if err != nil || input.GameID == "" || input.CardID == "" {
		log.Printf("error decoding values: %v", err)
		// TODO serve error message
		return
	}
	gameInternal, err := c.Svc.MarkCardUsed(input.GameID, input.CardID)
	if err != nil {
		log.Printf("error marking card %s as used: %v", input.CardID, err)
		// TODO server error message
		return
	}

	game := internalToExternal(gameInternal)
	gameData, err := json.Marshal(game)
	if err != nil {
		log.Printf("error marshalling game struct to json %v: %v", game, err)
		// TODO serve error message
		return
	}
	message := Envelope{
		ClientID: input.GameID,
		Message:  Message{Name: "gameState", Data: string(gameData)},
	}
	cl.pool.Broadcast <- message
}

// StartRound is controller to start a new round of the game by setting all cards to un-used state
func (c *controller) StartRound(cl *Client, data interface{}) {
	game := GameInput{}

	err := json.Unmarshal([]byte(data.(string)), &game)
	if err != nil || game.ID == "" {
		log.Printf("error decoding values: %v", err)
		// TODO serve error message
		return
	}

	newRound, err := c.Svc.StartRound(game.ID)
	if err != nil {
		log.Printf("error setting cards unused for game %s: %v", game.ID, err)
		// TODO serve error message
		return
	}

	gameData, err := json.Marshal(internalToExternal(newRound))
	if err != nil {
		log.Printf("error marshalling game struct to json %v: %v", newRound, err)
		// TODO serve error message
		return
	}
	message := Envelope{
		ClientID: game.ID,
		Message:  Message{Name: "gameState", Data: string(gameData)},
	}
	cl.pool.Broadcast <- message
}

// ResetGame is controller to delete all cards for a given game
func (c *controller) ResetGame(cl *Client, data interface{}) {
	game := GameInput{}

	err := json.Unmarshal([]byte(data.(string)), &game)
	if err != nil || game.ID == "" {
		log.Printf("error decoding values: %v", err)
		// TODO serve error message
		return
	}

	err = c.Svc.DeleteCards(game.ID)
	if err != nil {
		log.Printf("error deleting cards for game %s: %v", game.ID, err)
		// TODO serve error messaage
		return
	}

	gameInternal, err := c.Svc.GetGame(game.ID)
	if err != nil {
		log.Printf("error fetching cards: %v", err)
		// TODO serve error message
		return
	}

	gameData, err := json.Marshal(internalToExternal(gameInternal))
	if err != nil {
		log.Printf("error marshalling game struct to json %v: %v", gameInternal, err)
		// TODO serve error message
		return
	}
	message := Envelope{
		ClientID: game.ID,
		Message:  Message{Name: "gameState", Data: string(gameData)},
	}
	cl.pool.Broadcast <- message
}
