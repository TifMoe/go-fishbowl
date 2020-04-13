package api

import (
    "encoding/json"
    "fmt"
	"io/ioutil"	
	"path/filepath"
	"math/rand"
	"time"
)

// RandomService is interface to find random words for namespacing new games
type RandomService interface {
	GetRandomWords() (string, error)
}

// GetRandomWords is service to 
func GetRandomWords() (string, error) {
	words, err := readWords()
	if err != nil {
		return "", err
	}

	noun := randPicker(words.Nouns)
	adj := randPicker(words.Adjectives)

	return adj + "-" + noun, nil
}

func randPicker(words []string) string {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	return words[rand.Intn(len(words))]
}

func readWords() (words *WordChoices, err error) {
	// read file
	absPath, _ := filepath.Abs("./assets/randwords.json")  // path from the working directory
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