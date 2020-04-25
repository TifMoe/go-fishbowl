package api

import (
	"fmt"
	"log"
    "net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"github.com/tifmoe/go-fishbowl/src/errors"
	"github.com/tifmoe/go-fishbowl/src/service"
)

// NewGameController will instantiate a new game handler
func NewGameController(svc service.GameService) *controller {
	return &controller{
		Svc: svc,
	}
}

// GameController is interface with methods to interact with redis db	
type GameController interface {
	NewGame(w http.ResponseWriter, r *http.Request)
	GetGame(w http.ResponseWriter, r *http.Request)
	NewCard(w http.ResponseWriter, r *http.Request)
}

// Controller holds services and validators for Game Handlers
type controller struct {
	Svc service.GameService
}

// NewGame is controller for generating new game namespace and instantiating in the database
func (c *controller) NewGame(w http.ResponseWriter, r *http.Request) {

	// Instantiate new game namespace
	nameSpace, err := c.Svc.NewGame()
	if err != nil {
		log.Printf("error generating new namespace: %v", err)

		// Build and return error
		res := buildResponse(Game{}, errors.ErrNewGame, nameSpace)
		serveResponse(w, res)
		return
	}

	game := Game{
		ID: nameSpace,
	}
	res := buildResponse(game, &errors.ErrorInternal{}, nameSpace)

	// Return response
	serveResponse(w, res)
}

// NewCard is controller for saving a new card to the existing game
func (c *controller) NewCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gameID := params["gameID"]
	card := service.CardInput{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&card)
	if err != nil || card.IsEmpty() {
		log.Printf("error decoding values: %v", err)
		res := buildResponse(Game{}, errors.ErrInvalidInput, gameID)
		serveResponse(w, res)
		return
	}

	cardID, err := c.Svc.SaveCard(gameID, &card)
	if err != nil {
		log.Printf("error saving values: %v", err)
		// TODO discern validation and not-found errors
		res := buildResponse(Game{}, errors.ErrInternalError, gameID)
		serveResponse(w, res)
		return
	}

	game := Game{
		ID: gameID,
		Cards: []Card{
			Card{
				ID: cardID,
				Value: card.Value,
				Used: false,
			},
		},
	}
	msg := fmt.Sprintf("Successfully saved new card to %s", gameID)

	res := buildResponse(game, &errors.ErrorInternal{}, msg)
	res.Status = 201
	serveResponse(w, res)
}

// GetGame is controller for fetching a specific game
func (c *controller) GetGame(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gameID := params["gameID"]
	
	gameInternal, err := c.Svc.GetGame(gameID)
	if err != nil {
		log.Printf("error fetching cards: %v", err)
		res := buildResponse(Game{}, errors.ErrInternalError, gameID)
		serveResponse(w, res)
		return
	}
	game := internalToExternal(gameInternal)
	msg := fmt.Sprintf("Game %s has %d cards", gameID, len(game.Cards))

	res := buildResponse(game, &errors.ErrorInternal{}, msg)
	serveResponse(w, res)
}
