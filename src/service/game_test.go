package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tifmoe/go-fishbowl/src/repository"
)


func TestNewGameService(t *testing.T) {

	type test struct {
		name 				string
		mockRepo 			*mockRepo
		mockRand			*mockRand
		expectedResponse  	string
		expectedError		error
	}
	
	tests := []test{
        {
			name: "request to create new game succeeds",
			mockRepo: newMockRepo("hungry-hippo"), 
			mockRand: newMockRand("hungry-hippo"),
			expectedResponse: "hungry-hippo",
			expectedError: nil,
		},
		{
			name: "request to create new game fails on randomizer error",
			mockRepo: newMockRepo("hungry-hippo"), 
			mockRand: newMockRand("error"),
			expectedResponse: "",
			expectedError: fmt.Errorf("failed to generate random name"),
		},
		{
			name: "request to create new game fails on repository error",
			mockRepo: newMockRepo("error"), 
			mockRand: newMockRand("hungry-hippo"),
			expectedResponse: "",
			expectedError: fmt.Errorf("failed to save new game hungry-hippo"),
		},
    }

    for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := NewGameService(tc.mockRepo, tc.mockRand, 3)

			team1, team2 := "boo", "hoo"
			input := &TeamInput{Team1: &team1, Team2: &team2}
			gameID, err := svc.NewGame(input)
			assert.Equal(t, tc.expectedResponse, gameID)
			assert.Equal(t, tc.expectedError, err)
		})
    }
}

func TestSaveCardService(t *testing.T) {

	type test struct {
		name 				string
		newCard				*CardInput
		gameID				string
		maxCards			int
		expectedError		error
	}
	
	tests := []test{
        {
			name: "request to save new card succeeds",
			newCard: &CardInput{Value: "new-value"},
			gameID: "two-cards",
			maxCards: 30,
			expectedError: nil,
		},
		{
			name: "request to save new card fails if max cards already exist",
			newCard: &CardInput{Value: "new-value"},
			gameID: "two-cards",
			maxCards: 2,
			expectedError: fmt.Errorf("game already has maximum number of cards"),
		},
		{
			name: "request to save new card fails if game does not exist",
			newCard: &CardInput{Value: "new-value"},
			gameID: "error",
			maxCards: 10,
			expectedError: fmt.Errorf("game error does not exist"),
		},
		{
			name: "request to save new card fails if value not provided",
			newCard: &CardInput{},
			gameID: "two-cards",
			maxCards: 10,
			expectedError: fmt.Errorf("invalid card input"),
		},
		{
			name: "request to save new card fails if value exceeds max char number",
			newCard: &CardInput{Value: "exceedingly verbose individual enters in a card value which is far too long"},
			gameID: "two-cards",
			maxCards: 10,
			expectedError: fmt.Errorf("invalid card input"),
		},
		{
			name: "request to save new card fails if value exceeds max char number",
			newCard: &CardInput{Value: "exceedingly verbose individual enters in a card value which is far too long"},
			gameID: "two-cards",
			maxCards: 10,
			expectedError: fmt.Errorf("invalid card input"),
		},
		{
			name: "request to save new card fails if game could not be updated",
			newCard: &CardInput{Value: "valid-input"},
			gameID: "update-error",
			maxCards: 10,
			expectedError: fmt.Errorf("error updateing game with new card"),
		},
    }

    for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := newMockRepo(tc.gameID) 
			mockRand := newMockRand(tc.gameID)
			svc := NewGameService(mockRepo, mockRand, tc.maxCards)

			cardID, err := svc.SaveCard(tc.gameID, tc.newCard)
			assert.NotNil(t, cardID)
			assert.Equal(t, tc.expectedError, err)
		})
    }
}

func TestGetGameService(t *testing.T) {

	type test struct {
		name 				string
		gameID				string
		expectGame			*Game
		expectedError		error
	}
	
	tests := []test{
        {
			name: "request to get existing game succeeds",
			gameID: "two-cards",
			expectGame: &Game{
				ID: "hungry-hippos",
				Cards: []Card{
					Card{
						ID: "card-1",
						Value: "value-1",
						Used: false,
					},
					Card{
						ID: "card-2",
						Value: "value-2",
						Used: false,
					},
				},
				UnusedCards: 2,
			},
			expectedError: nil,
		},
		{
			name: "request to get game fails when repo returns an error",
			gameID: "error",
			expectGame: nil,
			expectedError: fmt.Errorf("failed to fetch game"),
		},
    }

    for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := newMockRepo(tc.gameID) 
			mockRand := newMockRand(tc.gameID)
			svc := NewGameService(mockRepo, mockRand, 10)

			game, err := svc.GetGame(tc.gameID)
			assert.Equal(t, tc.expectGame, game)
			assert.Equal(t, tc.expectedError, err)
		})
    }
}

func TestGetRandomCardService(t *testing.T) {

	type test struct {
		name 				string
		gameID				string
		mockRandCase		string
		expectCard			*Card
		expectedError		error
	}
	
	tests := []test{
        {
			name: "request to get random card succeeds",
			gameID: "two-cards",
			mockRandCase: "first", 
			expectCard: &Card{
				ID: "card-1",
				Value: "value-1",
				Used: false,
			},
			expectedError: nil,
		},
		{
			name: "request to get random card returns nil when all cards used",
			gameID: "all-cards-used",
			mockRandCase: "first", 
			expectCard: nil,
			expectedError: nil,
		},
    }

    for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := newMockRepo(tc.gameID) 
			mockRand := newMockRand(tc.mockRandCase)
			svc := NewGameService(mockRepo, mockRand, 10)

			card, err := svc.GetRandomCard(tc.gameID)
			assert.Equal(t, tc.expectCard, card)
			assert.Equal(t, tc.expectedError, err)
		})
    }
}

func TestMarkCardUsedService(t *testing.T) {

	type test struct {
		name 				string
		gameID				string
		cardID				string
		expectedError		error
	}
	
	tests := []test{
        {
			name: "request to mark card as used succeeds",
			gameID: "two-cards",
			cardID: "card-1",
			expectedError: nil,
		},
		{
			name: "request to mark card as used fails when card not found in game",
			gameID: "two-cards",
			cardID: "card-not-exists",
			expectedError: fmt.Errorf("card card-not-exists not found in game two-cards"),
		},
    }

    for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := newMockRepo(tc.gameID) 
			mockRand := newMockRand(tc.gameID)
			svc := NewGameService(mockRepo, mockRand, 10)

			err := svc.MarkCardUsed(tc.gameID, tc.cardID)
			assert.Equal(t, tc.expectedError, err)
		})
    }
}

func TestResetGameService(t *testing.T) {

	type test struct {
		name 				string
		gameID				string
		expectGame			*Game
		expectedError		error
	}
	
	tests := []test{
        {
			name: "request to set cards as unused succeeds",
			gameID: "all-cards-used",
			expectGame: &Game{
				ID: "hungry-hippos",
				Cards: []Card{
					Card{
						ID: "card-1",
						Value: "value-1",
						Used: false,
					},
				},
				Started: true,
				Round: 1,
				Team1Turn: true,
				UnusedCards: 1,
			},
			expectedError: nil,
		},
		{
			name: "request to set cards as unused fails when repository error getting game",
			gameID: "error",
			expectGame: nil,
			expectedError: fmt.Errorf("Game does not exist"),
		},
    }

    for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := newMockRepo(tc.gameID) 
			mockRand := newMockRand(tc.gameID)
			svc := NewGameService(mockRepo, mockRand, 10)

			game, err := svc.ResetGame(tc.gameID)
			assert.Equal(t, tc.expectGame, game)
			assert.Equal(t, tc.expectedError, err)
		})
    }
}


// Mock Repository Interface
func newMockRepo(c string) *mockRepo {
	return &mockRepo{
		Case: c,
	}
}

type mockRepo struct {
	Case string
}

func (s *mockRepo) SaveNewGame(gameID, team1, team2 string) error {
	switch s.Case {
	case "error":
		return fmt.Errorf("Error saving new game")
	default:
		return nil
	}
}

func (s *mockRepo) UpdateGame(game *repository.Game) error {
	switch s.Case {
	case "update-error":
		return fmt.Errorf("failed to update game")
	default:
		return nil
	}
}

func (s *mockRepo) GetGame(gameID string) (game *repository.Game, err error) {
	switch s.Case {
	case "error":
		return &repository.Game{}, fmt.Errorf("Game does not exist")
	case "two-cards":
		return &repository.Game{
			ID: "hungry-hippos",
			Cards: []repository.Card{
				repository.Card{
					ID: "card-1",
					Value: "value-1",
					Used: false,
				},
				repository.Card{
					ID: "card-2",
					Value: "value-2",
					Used: false,
				},
			},
		}, nil
	case "all-cards-used":
		return &repository.Game{
			ID: "hungry-hippos",
			Cards: []repository.Card{
				repository.Card{
					ID: "card-1",
					Value: "value-1",
					Used: true,
				},
			},
		}, nil
	default:
		return &repository.Game{}, nil
	}
}

// Mock Random Interface
func newMockRand(c string) *mockRand {
	return &mockRand{
		Value: c,
	}
}

type mockRand struct {
	Value string
}

func (s *mockRand) GetRandomWords() (string, error) {
	switch s.Value {
	case "error":
		return "", fmt.Errorf("Error generating new word")
	default:
		return s.Value, nil
	}
}

func (s *mockRand) GetRandomCard(cards []Card) *Card {
	switch s.Value {
	case "first":
		return &Card{
			ID: "card-1",
			Value: "value-1",
			Used: false,
		}
	default:
		return &Card{}
	}
}