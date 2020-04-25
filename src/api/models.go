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
	ID     string 	`json:"id"`
	Cards   []Card	`json:"cards"`
}

// Card contains data for a specific card
type Card struct {
	ID 		string	`json:"id"`
	Value   string  `json:"value"`
	Used 	bool  	`json:"used"`
}

func internalToExternal(g *service.Game) Game {
	game := Game{}
	game.ID = g.ID

	cardCount := len(g.Cards)
	cards := make([]Card, 0, cardCount)
	if cardCount > 0 {
		for _, card := range g.Cards {
			c := Card{}
			c.ID = card.ID
			c.Value = card.Value
			c.Used = card.Used
			cards = append(cards, c)
		}
	}
	game.Cards = cards
	return game
}
