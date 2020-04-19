package api

import (
	"fmt"
	"log"
    "net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"github.com/tifmoe/go-fishbowl/src/service"
)

// NewGameController will instantiate a new game handler
func NewGameController(svc *service.Service) *Controller {
	return &Controller{
		Svc: svc,
	}
}

// GameController is interface with methods to interact with redis db	
type GameController interface {
    NewGame(w http.ResponseWriter, r *http.Request)
	SaveNewCard(w http.ResponseWriter, r *http.Request)
	FetchCards(w http.ResponseWriter, r *http.Request)
}

// Controller holds services and validators for Game Handlers
type Controller struct {
	Svc *service.Service
}

// NewGame is controller for generating new game namespace and instantiating in the database
func (c *Controller) NewGame(w http.ResponseWriter, r *http.Request) {

	// Instantiate new game namespace
	nameSpace, err := c.Svc.NewGame()
	if err != nil {
		log.Printf("error generating new namespace: %v", err)
		// ToDo generate error response here
		return
	}

	res := &Response{
		Message: nameSpace,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// SaveNewCard is controller for saving a new card to the existing game
func (c *Controller) SaveNewCard(w http.ResponseWriter, r *http.Request) {
	// Fetch game ID from request path
	params := mux.Vars(r)
	gameID := params["gameID"]
	card := service.CardInput{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&card)
	if err != nil {
		log.Printf("error decoding values: %v", err)
		return // TODO: Return 400
	}

	err = c.Svc.SaveNewCard(gameID, &card)
	if err != nil {
		log.Printf("error saving values: %v", err)
		return // TODO: Return 500
	}

	res := &Response{
		Message: fmt.Sprintf("Successfully saved to %s", gameID),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// FetchCards is controller for saving a new card to the existing game
func (c *Controller) FetchCards(w http.ResponseWriter, r *http.Request) {
	// Get game ID from request path
	params := mux.Vars(r)
	gameID := params["gameID"]
	
	cards, err := c.Svc.FetchCards(gameID)
	if err != nil {
		log.Printf("error fetching cards: %v", err)
		return // TODO Return 500
	}
	res := &Response{
		Message: fmt.Sprintf("Game %s identified! Here are the cards: %v", gameID, cards),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
