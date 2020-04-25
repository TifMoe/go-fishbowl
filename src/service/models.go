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
	Value		string `json:"value" validate:"required,min=2,max=50"`
}

// IsEmpty will return true if Error struct does not contain something other than default
func (c CardInput) IsEmpty() bool {
    return c.Value == ""
}

// Game is the internal struct for a game object
type Game struct {
	ID     string 		`json:"id,omitempty"`
	Cards   []Card		`json:"cards,omitempty"`
}

// Card is the internal struct for a card object
type Card struct {
	ID 		string		`json:"id,omitempty"`
	Value   string     `json:"value,omitempty"`
	Used 	bool       `json:"used"`
}

func gameDTOtoInternal(dto *repository.Game) *Game {
	game := &Game{}
	game.ID = dto.ID

	cardCount := len(dto.Cards)
	cards := make([]Card, 0, cardCount)
	if cardCount > 0 {
		for _, card := range dto.Cards {
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
