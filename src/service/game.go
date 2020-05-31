package service

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"gopkg.in/go-playground/validator.v9"

	"github.com/tifmoe/go-fishbowl/src/repository"
)

// NewGameService will instantiate a new game service
func NewGameService(r repository.Repository, rand RandomService, max int) GameService {
	v := validator.New()
	return &service{
		Repo:         r,
		Validate:     v,
		MaximumCards: max,
		Rand:         rand,
	}
}

// GameService is interface with methods to interact with redis db
type GameService interface {
	NewGame(input *TeamInput) (string, error)
	GetGame(gameID string) (game *Game, err error)
	UpdateGame(gameID string, input *GameInput) (*Game, error)

	SaveCard(gameID string, card *CardInput) (int, error)
	GetRandomCard(gameID string) (card *Card, err error)
	MarkCardUsed(gameID, cardID string) (*Game, error)
	StartRound(gameID string) (*Game, error)
	DeleteCards(gameID string) error
}

type service struct {
	Repo         repository.Repository
	Validate     *validator.Validate
	MaximumCards int
	Rand         RandomService
}

// NewGame is service for generating new game namespace and instantiating in the database
func (s *service) NewGame(input *TeamInput) (string, error) {
	// Terminate the request if the input is not valid
	if err := s.Validate.Struct(input); err != nil {
		log.Printf("error validating values from game input %v: %v", input, err)
		return "", fmt.Errorf("invalid game input")
	}

	// Generate new random word pair for namespace
	nameSpace, err := s.Rand.GetRandomWords()
	if err != nil {
		log.Printf("error generating random name: %v", err)
		return "", fmt.Errorf("failed to generate random name")
	}

	// Attempt to save to database
	err = s.Repo.SaveNewGame(nameSpace, *input.Team1, *input.Team2)
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

	// Terminate the request if the input is not valid
	if err := s.Validate.Struct(input); err != nil {
		log.Printf("error validating values from game input %v: %v", input, err)
		return game, fmt.Errorf("invalid game input")
	}

	existingGame, err := s.Repo.GetGame(gameID)
	if existingGame == nil || err != nil {
		log.Printf("attempted to update a non-existent game %s: %v", gameID, err)
		return game, fmt.Errorf("game %s does not exist", gameID)
	}

	// TODO: There's a better way to handle partial updates
	if input.Started != nil {
		existingGame.Started = *input.Started
	}
	if input.Round != nil {
		existingGame.Round = *input.Round
	}
	if input.Team1Turn != nil {
		existingGame.Team1Turn = *input.Team1Turn
	}

	err = s.Repo.UpdateGame(existingGame)
	if err != nil {
		log.Printf("error updating game %v: %v", gameID, err)
		return game, fmt.Errorf("error updating game")
	}
	return gameDTOtoInternal(existingGame), nil
}

// SaveCard is controller for saving a new card to the existing game and returns count of cards in game
func (s *service) SaveCard(gameID string, card *CardInput) (int, error) {

	// Terminate the request if the input is not valid
	if err := s.Validate.Struct(card); err != nil {
		log.Printf("error validating values from card %v: %v", card, err)
		return 0, fmt.Errorf("invalid card input")
	}

	existingGame, err := s.Repo.GetGame(gameID)
	if existingGame == nil || err != nil {
		log.Printf("attempted to save card to non-existent game %s: %v", gameID, err)
		return 0, fmt.Errorf("game %s does not exist", gameID)
	}

	if len(existingGame.Cards) >= s.MaximumCards {
		log.Printf("maximum cards already added for game %s: %d", gameID, s.MaximumCards)
		return 0, fmt.Errorf("game already has maximum number of cards")
	}

	newCard := &repository.Card{
		ID:    uuid.New().String(),
		Value: card.Value,
		Used:  false,
	}
	existingGame.Cards = existingGame.AddCard(*newCard)

	err = s.Repo.UpdateGame(existingGame)
	if err != nil {
		log.Printf("error saving card %v: %v", card, err)
		return 0, fmt.Errorf("error updateing game with new card")
	}
	return len(existingGame.Cards), nil
}

// GetGame is service to fetch the requested game
// TODO support request parameters to fetch specific resources only
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

// GetRandomCard is service to draw a new unused card
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

// MarkCardUsed is service to update existing card to used and record the team responsible for the current round
func (s *service) MarkCardUsed(gameID, cardID string) (*Game, error) {
	var game *Game

	gameDTO, err := s.Repo.GetGame(gameID)
	if err != nil {
		log.Printf("error fetching game %v: %v", gameID, err)
		return game, err
	}

	var currentTeam *repository.Team
	if gameDTO.Team1Turn {
		currentTeam = &gameDTO.Teams.Team1
	} else {
		currentTeam = &gameDTO.Teams.Team2
	}
	currentRound := gameDTO.Round

	found := false
	for i := range gameDTO.Cards {
		if gameDTO.Cards[i].ID == cardID {
			gameDTO.Cards[i].Used = true
			currentTeam.IncrementPoints(currentRound, &gameDTO.Cards[i])
			found = true
		}
	}

	if !found {
		return game, fmt.Errorf("card %s not found in game %s", cardID, gameID)
	}

	err = s.Repo.UpdateGame(gameDTO)
	if err != nil {
		log.Printf("error updating card %v: %v", cardID, err)
		return game, err
	}
	game = gameDTOtoInternal(gameDTO)
	return game, nil
}

// StartRound is service to reset game to default values including moving all cards to un-used state before starting fresh round
func (s *service) StartRound(gameID string) (*Game, error) {
	var game *Game

	gameDTO, err := s.Repo.GetGame(gameID)
	if err != nil {
		log.Printf("error fetching game %v: %v", gameID, err)
		return game, err
	}

	for i := range gameDTO.Cards {
		gameDTO.Cards[i].Used = false
	}

	// Set default values for new round
	gameDTO.Started = true
	gameDTO.Round = gameDTO.Round + 1
	gameDTO.Team1Turn = !(gameDTO.Round%2 == 0) // Team 1 should start for every odd round

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
	gameDTO.Round = 0
	gameDTO.Team1Turn = true

	err = s.Repo.UpdateGame(gameDTO)
	if err != nil {
		log.Printf("error updating game %v: %v", gameID, err)
		return err
	}
	return nil
}
