package api

// Response contains message in json response
type Response struct {
	Message		string `json:"message"`
}

// WordChoices is struct to hold selection of random words to choose from
type WordChoices struct {
	Nouns 	   []string 	`json:"nouns"`
	Adjectives []string 	`json:"adjectives"`
}