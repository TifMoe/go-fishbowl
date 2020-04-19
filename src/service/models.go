package service

// WordChoices is struct to hold selection of random words to choose from
type WordChoices struct {
	Nouns 	   []string 	`json:"nouns"`
	Adjectives []string 	`json:"adjectives"`
}

// CardInput contains value for new card from game
type CardInput struct {
	Value		string `json:"value" validate:"required,min=2,max=50"`
}
