package service

import (
	"github.com/tifmoe/go-fishbowl/src/repository"
)

// WordChoices is struct to hold selection of random words to choose from
type WordChoices struct {
	Nouns      []string `json:"nouns"`
	Adjectives []string `json:"adjectives"`
}

// CardInput contains value for new card from game
type CardInput struct {
	GameID string `json:"gameID" validate:"required"`
	Value  string `json:"value" validate:"required,min=2,max=30"`
}

// IsEmpty will return true if Error struct does not contain something other than default
func (c CardInput) IsEmpty() bool {
	return c.Value == ""
}

// GameInput contains value for updates to game data
type GameInput struct {
	ID        string `json:"gameID" validate:"required"`
	Started   *bool  `json:"started,omitempty"`
	Round     *int   `json:"current_round,omitempty" validate:"max=5"`
	Team1Turn *bool  `json:"team_1_turn,omitempty"`
}

// TeamInput contains value for team names
type TeamInput struct {
	Team1 *string `json:"team_1,omitempty" validate:"required,min=2,max=30"`
	Team2 *string `json:"team_2,omitempty" validate:"required,min=2,max=30"`
}

// IsEmpty will return true if Error struct does not contain something other than default
func (t TeamInput) IsEmpty() bool {
	return &t.Team1 == nil || &t.Team2 == nil
}

// Game is the internal struct for a game object
type Game struct {
	ID          string `json:"id,omitempty"`
	Cards       []Card `json:"cards,omitempty"`
	Started     bool   `json:"started,omitempty"`
	Round       int    `json:"current_round,omitempty"`
	Team1Turn   bool   `json:"team_1_turn,omitempty"`
	UnusedCards int    `json:"unused_cards,omitempty"`
	Teams       Teams  `json:"teams,omitempty"`
}

// Card is the internal struct for a card object
type Card struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
	Used  bool   `json:"used,omitempty"`
}

// Teams is nested struct containing details for each of the two teams
type Teams struct {
	Team1 Team `json:"team_1,omitempty"`
	Team2 Team `json:"team_2,omitempty"`
}

// Team contains data for a specific
type Team struct {
	Name   string   `json:"name,omitempty"`
	Round1 []string `json:"round_1,omitempty"`
	Round2 []string `json:"round_2,omitempty"`
	Round3 []string `json:"round_3,omitempty"`
	Round4 []string `json:"round_4,omitempty"`
}

// TODO: This is gross, optimize later
func gameDTOtoInternal(dto *repository.Game) *Game {
	game := &Game{}
	game.ID = dto.ID
	game.Started = dto.Started
	game.Round = dto.Round
	game.Team1Turn = dto.Team1Turn

	// update teams
	game.Teams.Team1 = dtoToInternalTeam(dto.Teams.Team1)
	game.Teams.Team2 = dtoToInternalTeam(dto.Teams.Team2)

	unusedCount := 0
	cardCount := len(dto.Cards)
	cards := make([]Card, 0, cardCount)
	if cardCount > 0 {
		for _, card := range dto.Cards {
			if !card.Used {
				unusedCount = unusedCount + 1
			}
			c := Card{}
			c.ID = card.ID
			c.Value = card.Value
			c.Used = card.Used
			cards = append(cards, c)
		}
	}
	game.Cards = cards
	game.UnusedCards = unusedCount

	return game
}

func dtoToInternalTeam(dto repository.Team) Team {
	t := Team{}
	t.Name = dto.Name
	t.Round1 = dto.Round1
	t.Round2 = dto.Round2
	t.Round3 = dto.Round3
	t.Round4 = dto.Round4
	return t
}
