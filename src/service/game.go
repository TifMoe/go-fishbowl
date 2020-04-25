package service

import (
	"fmt"
	"log"

	"gopkg.in/go-playground/validator.v9"
	"github.com/google/uuid"

	"github.com/tifmoe/go-fishbowl/src/repository"
)

// NewGameService will instantiate a new game service
func NewGameService(r repository.Repository, v *validator.Validate, max int) *service {
	return &service{
		Repo: r,
		Validate: v,
		MaximumCards: max,
	}
}

// GameService is interface with methods to interact with redis db	
type GameService interface {
    NewGame() (string, error)
	SaveCard(gameID string, card *CardInput) (string, error)
	GetGame(gameID string) (game *Game, err error)
}

type service struct {
	Repo 			repository.Repository
	Validate 		*validator.Validate
	MaximumCards 	int
}

// NewGame is service for generating new game namespace and instantiating in the database
func (s *service) NewGame() (string, error) {
	// Generate new random word pair for namespace
	nameSpace, err := GetRandomWords()
	if err != nil {
		return "", err
	}

	// Attempt to save to database
	err = s.Repo.SaveNewGame(nameSpace)
	if err != nil {
		return "", err
	}

	// ToDo, if name exists, generate new one

	return nameSpace, nil
}

// SaveCard is controller for saving a new card to the existing game
func (s *service) SaveCard(gameID string, card *CardInput) (string, error) {

	// Terminate the request if the input is not valid
	if err := s.Validate.Struct(card); err != nil {
		log.Printf("error validating values from card %v: %v", card, err)
		return "", err
	}

	existingGame, err := s.Repo.GetGame(gameID)

	if existingGame == nil || err != nil{
		fmt.Printf("Attempted to save card to non-existent game %s: %v\n", gameID, err)
		return "", err
	} 

	if len(existingGame.Cards) >= s.MaximumCards {
		err = fmt.Errorf("game already has maximum number of cards")
		return "", err
	}

	newCard := &repository.Card{
		ID: uuid.New().String(),
		Value: card.Value, 
		Used: false,
	}
	existingGame.Cards = existingGame.AddCard(*newCard)

	err = s.Repo.UpdateGame(existingGame)
	if err != nil {
		log.Printf("error saving card %v: %v", card, err)
		return "", err
	}
	return newCard.ID, nil
}

// GetGame is controller for saving a new card to the existing game
func (s *service) GetGame(gameID string) (game *Game, err error) {
	gameDTO, err := s.Repo.GetGame(gameID)
	fmt.Printf("Here is the game: %+v", gameDTO)
	if err != nil {
		return
	}
	game = gameDTOtoInternal(gameDTO)
	fmt.Printf("here is the gaame internal: %+v", game)
	return
}
