package api

import (
	"github.com/tifmoe/go-fishbowl/src/errors"
	"github.com/tifmoe/go-fishbowl/src/service"
)

// apiResponse contains status code and json response to be served by API
type apiResponse struct {
	Status 	int 
	Data 	*Response
}

// Response contains json response
type Response struct {
	Result 		[]Game 			`json:"result"`
	Success		bool 			`json:"success"`
	Error 		[]errors.Error 	`json:"error"`
	Message		string 			`json:"message"`
}

// Game contains data a specific game session
type Game struct {
	ID     		string 	`json:"id"`
	Cards   	[]Card	`json:"cards"`
	Started		bool 	`json:"started"`
	Round		int		`json:"current_round"`
	UnusedCards	int		`json:"unused_cards"`
}

// Card contains data for a specific card
type Card struct {
	ID 		string	`json:"id"`
	Value   string  `json:"value"`
	Used 	bool  	`json:"used"`
}

// TODO: This is gross, optimize later
func internalToExternal(g *service.Game) Game {
	game := Game{}
	game.ID = g.ID
	game.Started = g.Started
	game.Round = g.Round
	game.UnusedCards = g.UnusedCards

	cardCount := len(g.Cards)
	cards := make([]Card, 0, cardCount)
	if cardCount > 0 {
		for _, card := range g.Cards {
			c := internalToExternalCard(&card)
			cards = append(cards, c)
		}
	}
	game.Cards = cards
	return game
}

func internalToExternalCard(card *service.Card) Card {
	c := Card{}
	c.ID = card.ID
	c.Value = card.Value
	c.Used = card.Used
	return c
}