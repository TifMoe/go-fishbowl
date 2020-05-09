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
	ID     			string 	`json:"id"`
	Cards   		[]Card	`json:"cards"`
	Started			bool 	`json:"started"`
	CurrentRound	int		`json:"current_round"`
	Team1Turn		bool	`json:"team_1_turn"`
	UnusedCards		int		`json:"unused_cards"`
	Teams			Teams 	`json:"teams"`
}

// Card contains data for a specific card
type Card struct {
	ID 		string	`json:"id"`
	Value   string  `json:"value"`
	Used 	bool  	`json:"used"`
}

// Teams is nested struct containing details for each of the two teams
type Teams struct {
	Team1 	Team `json:"team_1"`
	Team2 	Team `json:"team_2"`
}

// Team contains data for a specific
type Team struct {
	Name		string	`json:"name"`
	Round1		int		`json:"round_1_pts"`
	Round2   	int		`json:"round_2_pts"`
	Round3   	int		`json:"round_3_pts"`
	Round4   	int		`json:"round_4_pts"`
}

// TODO: This is gross, optimize later
func internalToExternal(g *service.Game) Game {
	game := Game{}
	game.ID = g.ID
	game.Started = g.Started
	game.CurrentRound = g.Round
	game.Team1Turn = g.Team1Turn
	game.UnusedCards = g.UnusedCards

	// update teams
	game.Teams.Team1 = internalToExternalTeam(&g.Teams.Team1)
	game.Teams.Team2 = internalToExternalTeam(&g.Teams.Team2)

	// update cards
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

func internalToExternalTeam(team *service.Team) Team {
	t := Team{}
	t.Name = team.Name
	t.Round1 = len(team.Round1)
	t.Round2 = len(team.Round2)
	t.Round3 = len(team.Round3)
	t.Round4 = len(team.Round4)
	return t
}