package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"time"
)

// NewRandomService will instantiate a new random service
func NewRandomService() RandomService {
	return &random{}
}

// RandomService is interface for collection of methods to draw random values
type RandomService interface {
	GetRandomCard(cards []Card) *Card
	GetRandomWords() (string, error)
}

type random struct{}

// GetRandomCard is utility function to select a random card from a slice
func (r *random) GetRandomCard(cards []Card) *Card {
	rand.Seed(time.Now().Unix())
	return &cards[rand.Intn(len(cards))]
}

// GetRandomWords is utility function to pick a random pairing of adj + noun to make game namespace
func (r *random) GetRandomWords() (string, error) {
	words, err := readWords()
	if err != nil {
		return "", err
	}

	noun := randPicker(words.Nouns)
	adj := randPicker(words.Adjectives)

	return adj + "-" + noun, nil
}

func randPicker(words []string) string {
	rand.Seed(time.Now().Unix())
	return words[rand.Intn(len(words))]
}

func readWords() (words *WordChoices, err error) {
	// TODO store words in noSQL db when we create it
	// read file from assets for now
	absPath, _ := filepath.Abs("./assets/randwords.json") // path from the working directory
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// unmarshall data into word choices struct
	err = json.Unmarshal(data, &words)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	return
}
