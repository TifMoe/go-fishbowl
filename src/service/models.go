package service

import (
	"github.com/tifmoe/go-fishbowl/src/repository"
)

// WordChoices is struct to hold selection of random words to choose from
type WordChoices struct {
	Nouns 	   []string 	`json:"nouns"`
	Adjectives []string 	`json:"adjectives"`
}

// CardInput contains value for new card from game
type CardInput struct {
	Value		string `json:"value" validate:"required,min=2,max=30"`
}

// IsEmpty will return true if Error struct does not contain something other than default
func (c CardInput) IsEmpty() bool {
    return c.Value == ""
}

// GameInput contains value for updates to game data
type GameInput struct {
	Started		bool 	`json:"started,omitempty"`
	Round		int		`json:"current_round,omitempty"`
}

// Game is the internal struct for a game object
type Game struct {
	ID     string 		`json:"id,omitempty"`
	Cards   []Card		`json:"cards,omitempty"`
	Started		bool 	`json:"started"`
	Round		int		`json:"current_round"`
	UnusedCards	int		`json:"unused_cards"`
}

// Card is the internal struct for a card object
type Card struct {
	ID 		string		`json:"id,omitempty"`
	Value   string     `json:"value,omitempty"`
	Used 	bool       `json:"used"`
}

// TODO: This is gross, optimize later
func gameDTOtoInternal(dto *repository.Game) *Game {
	game := &Game{}
	game.ID = dto.ID
	game.Started = dto.Started
	game.Round = dto.Round

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
