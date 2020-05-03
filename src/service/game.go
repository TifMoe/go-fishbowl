package service

import (
	"fmt"
	"log"

	"gopkg.in/go-playground/validator.v9"
	"github.com/google/uuid"

	"github.com/tifmoe/go-fishbowl/src/repository"
)

// NewGameService will instantiate a new game service
func NewGameService(r repository.Repository, rand RandomService, max int) GameService {
	v := validator.New()
	return &service{
		Repo: r,
		Validate: v,
		MaximumCards: max,
		Rand: rand,
	}
}

// GameService is interface with methods to interact with redis db	
type GameService interface {
    NewGame() (string, error)
	GetGame(gameID string) (game *Game, err error)
	UpdateGame(gameID string, input *GameInput) (*Game, error)

	SaveCard(gameID string, card *CardInput) (string, error)
	GetRandomCard(gameID string) (card *Card, err error)
	MarkCardUsed(gameID, cardID string) error
	SetCardsUnused(gameID string) (*Game, error)
	DeleteCards(gameID string) error
}

type service struct {
	Repo 			repository.Repository
	Validate 		*validator.Validate
	MaximumCards 		int
	Rand 			RandomService
}

// NewGame is service for generating new game namespace and instantiating in the database
func (s *service) NewGame() (string, error) {
	// Generate new random word pair for namespace
	nameSpace, err := s.Rand.GetRandomWords()
	if err != nil {
		log.Printf("error generating random name: %v", err)
		return "", fmt.Errorf("failed to generate random name")
	}

	// Attempt to save to database
	err = s.Repo.SaveNewGame(nameSpace)
	if err != nil {
		// TODO: if name is in play, automaticallly generate new one
		log.Printf("error generating new game: %v", err)
		return "", fmt.Errorf("failed to save new game %s", nameSpace)
	}

	return nameSpace, nil
}

// UpdateGame is service for updating a game
func (s *service) UpdateGame(gameID string, input *GameInput) (*Game, error) {
	var game *Game
	existingGame, err := s.Repo.GetGame(gameID)
	if existingGame == nil || err != nil{
		log.Printf("attempted to save card to non-existent game %s: %v", gameID, err)
		return game, fmt.Errorf("game %s does not exist", gameID)
	}

	existingGame.Started = input.Started
	existingGame.Round = input.Round

	err = s.Repo.UpdateGame(existingGame)
	if err != nil {
		log.Printf("error updating game %v: %v", gameID, err)
		return game, fmt.Errorf("error updating game")
	}
	return gameDTOtoInternal(existingGame), nil
}

// SaveCard is controller for saving a new card to the existing game
func (s *service) SaveCard(gameID string, card *CardInput) (string, error) {

	// Terminate the request if the input is not valid
	if err := s.Validate.Struct(card); err != nil {
		log.Printf("error validating values from card %v: %v", card, err)
		return "", fmt.Errorf("invalid card input")
	}

	existingGame, err := s.Repo.GetGame(gameID)
	if existingGame == nil || err != nil{
		log.Printf("attempted to save card to non-existent game %s: %v", gameID, err)
		return "", fmt.Errorf("game %s does not exist", gameID)
	} 

	if len(existingGame.Cards) >= s.MaximumCards {
		log.Printf("maximum cards already added for game %s: %d", gameID, s.MaximumCards)
		return "", fmt.Errorf("game already has maximum number of cards")
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
		return "", fmt.Errorf("error updateing game with new card")
	}
	return newCard.ID, nil
}

// GetGame is controller for saving a new card to the existing game
func (s *service) GetGame(gameID string) (*Game, error) {
	var game *Game
	gameDTO, err := s.Repo.GetGame(gameID)
	if err != nil {
		log.Printf("error fetching game %v: %v", gameID, err)
		return game, fmt.Errorf("failed to fetch game")
	}
	game = gameDTOtoInternal(gameDTO)
	return game, nil
}

// GetGame is controller for saving a new card to the existing game
func (s *service) GetRandomCard(gameID string) (card *Card, err error) {
	gameDTO, err := s.Repo.GetGame(gameID)
	if err != nil {
		return
	}
	unusedCards := []Card{}
	game := gameDTOtoInternal(gameDTO)
	for i := range game.Cards {
		if !game.Cards[i].Used {
			unusedCards = append(unusedCards, game.Cards[i])
		}
	}
	if len(unusedCards) == 0 {
		// Return nil card with no error if all cards used
		return
	}
	card = s.Rand.GetRandomCard(unusedCards)
	return
}

// MarkCardUsed is service to update existing card
func (s *service) MarkCardUsed(gameID, cardID string) error {
    gameDTO, err := s.Repo.GetGame(gameID)
    if err != nil {
        log.Printf("error fetching game %v: %v", gameID, err)
        return err
	}

	found := false
    for i := range gameDTO.Cards {
        if gameDTO.Cards[i].ID == cardID {
			gameDTO.Cards[i].Used = true
			found = true
        }
    }

	if !found {
		return fmt.Errorf("card %s not found in game %s", cardID, gameID)
	}

    err = s.Repo.UpdateGame(gameDTO)
    if err != nil {
        log.Printf("error updating card %v: %v", cardID, err)
        return err
    }
    return nil
}

// SetCardsUnused is service to set all cards for a game into an un-used state before starting fresh round
func (s *service) SetCardsUnused(gameID string) (*Game, error) {
	var game *Game

    gameDTO, err := s.Repo.GetGame(gameID)
    if err != nil {
        log.Printf("error fetching game %v: %v", gameID, err)
        return game, err
	}

    for i := range gameDTO.Cards {
        gameDTO.Cards[i].Used = false
    }

    err = s.Repo.UpdateGame(gameDTO)
    if err != nil {
        log.Printf("error updating game %v: %v", gameID, err)
        return game, err
	}
    return gameDTOtoInternal(gameDTO), nil
}

// DeleteCards is service to delete all cards for a given game
func (s *service) DeleteCards(gameID string) error {
	gameDTO, err := s.Repo.GetGame(gameID)
    if err != nil {
        log.Printf("error fetching game %v: %v", gameID, err)
        return err
	}
	gameDTO.Cards = []repository.Card{}

    err = s.Repo.UpdateGame(gameDTO)
    if err != nil {
        log.Printf("error updating game %v: %v", gameID, err)
        return err
	}
    return nil
}
