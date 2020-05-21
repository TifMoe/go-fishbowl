package repository

import "log"

// Game is the DTO for a specific game
type Game struct {
	ID        string `json:"id,omitempty"`
	Cards     []Card `json:"cards,omitempty"`
	Started   bool   `json:"started"`
	Round     int    `json:"current_round"`
	Team1Turn bool   `json:"team_1_turn"`
	Teams     Teams  `json:"teams"`
}

// Card is the DTO for card values
type Card struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
	Used  bool   `json:"used"`
}

// Teams is nested struct containing details for each of the two teams
type Teams struct {
	Team1 Team `json:"team_1"`
	Team2 Team `json:"team_2"`
}

// Team contains data for a specific
type Team struct {
	Name   string   `json:"name"`
	Round1 []string `json:"round_1"`
	Round2 []string `json:"round_2"`
	Round3 []string `json:"round_3"`
	Round4 []string `json:"round_4"`
}

// AddCard is helper function to add a new gard to an existing game
func (game *Game) AddCard(c Card) []Card {
	game.Cards = append(game.Cards, c)
	return game.Cards
}

// IncrementPoints is helper function to increment the number of points a team has won in a given round
func (team *Team) IncrementPoints(round int, card *Card) {
	switch round {
	case 1:
		team.Round1 = append(team.Round1, card.ID)
		return
	case 2:
		team.Round2 = append(team.Round2, card.ID)
		return
	case 3:
		team.Round3 = append(team.Round3, card.ID)
		return
	case 4:
		team.Round4 = append(team.Round4, card.ID)
		return
	default:
		log.Printf("current round invaid")
		return
	}
}
