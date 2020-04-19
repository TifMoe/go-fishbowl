package service

import (
	"fmt"
	"log"

	"gopkg.in/go-playground/validator.v9"

	"github.com/tifmoe/go-fishbowl/src/repository"
)

// NewGameService will instantiate a new game service
func NewGameService(r repository.Repository, v *validator.Validate) *Service {
	return &Service{
		Repo: r,
		Validate: v,
	}
}

// GameService is interface with methods to interact with redis db	
type GameService interface {
    NewGame() (string, error)
	SaveCard(gameID string, card *CardInput) error
	FetchCards(gameID string) (string, error)
}

// Service holds dependencies of service layer
type Service struct {
	Repo repository.Repository
	Validate *validator.Validate
}

// NewGame is service for generating new game namespace and instantiating in the database
func (s *Service) NewGame() (string, error) {
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

// SaveNewCard is controller for saving a new card to the existing game
func (s *Service) SaveNewCard(gameID string, card *CardInput) error {

	// Terminate the request if the input is not valid
	if err := s.Validate.Struct(card); err != nil {
		log.Printf("error validating values from card %v: %v", card, err)
		return err
	}

	err := s.Repo.SaveCard(gameID, card.Value)
	if err != nil {
		log.Printf("error saving card %v: %v", card, err)
		return err
	}
	return nil
}

// FetchCards is controller for saving a new card to the existing game
func (s *Service) FetchCards(gameID string) (cards string, err error) {
	game, err := s.Repo.GetGame(gameID)
	fmt.Printf("Here is the game: %+v", game)
	if err != nil {
		return 
	}
	return "Yes", nil
}
