package repository

// Game is the DTO for a specific game
type Game struct {
	ID     string 		`json:"id,omitempty"`
	Cards   []Card		`json:"cards,omitempty"`
}

// Card is the DTO for card values
type Card struct {
	ID 		string		`json:"id,omitempty"`
	Value   string     `json:"value,omitempty"`
	Used 	bool       `json:"used"`
}

// AddCard is helper function to add a new gard to an existing game
func (game *Game) AddCard(c Card) []Card {
    game.Cards = append(game.Cards, c)
    return game.Cards
}