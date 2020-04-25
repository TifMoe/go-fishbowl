package api

import (
	"fmt"
	"log"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"

	"github.com/stretchr/testify/assert"

	"github.com/tifmoe/go-fishbowl/src/errors"
	"github.com/tifmoe/go-fishbowl/src/service"
)


func TestNewGameController(t *testing.T) {

	type test struct {
		name 				string
        mockService 		*mockService
		expectedStatus		int
        expectedResponse  	Response
	}
	
	tests := []test{
        {
			name: "request to create new game succeeds",
			mockService: newMockService("successful-game"), 
			expectedStatus: 200, 
			expectedResponse: Response{
				Result: []Game{
					Game{
						ID: "successful-game",
					},
				},
				Success: true,
				Error: []errors.Error{},
				Message: "successful-game",
			},
		},
		{
			name: "request to create new game returns error",
			mockService: newMockService("error"), 
			expectedStatus: 500, 
			expectedResponse: Response{
				Result: []Game{},
				Success: false,
				Error: []errors.Error{
					errors.Error{
						Code: 1000,
						Message: "Failed to instantiate new game session, please try again later",
					},
				},
				Message: "",
			},
		},
    }

    for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var (
				response = Response{}
			)

			req, err := http.NewRequest("POST", "/v1/api/game", nil)
			assert.Nil(t, err)

			rr := httptest.NewRecorder()
			controller := NewGameController(tc.mockService)
			handler := http.HandlerFunc(controller.NewGame)
			handler.ServeHTTP(rr, req)


			if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
				log.Fatalln(err)
			}

			assert.Equal(t, tc.expectedStatus, rr.Code)
			assert.Equal(t, tc.expectedResponse, response)
		})
    }
}

func TestNewCardController(t *testing.T) {

	type test struct {
		name 				string
		mockService 		*mockService
		request 			func() *http.Request
		expectedStatus		int
        expectedResponse  	Response
	}
	
	tests := []test{
        {
			name: "request to save new card succeeds",
			mockService: newMockService("test-uuid"), 
			request: func() *http.Request {
				var data = []byte(`{"value":"Super funny new noun"}`)
				req, err := http.NewRequest("POST", "/v1/api/game/test-name/card", bytes.NewBuffer(data))
				assert.Nil(t, err)
				return req
			},
			expectedStatus: 201, 
			expectedResponse: Response{
				Result: []Game{
					Game{
						ID: "test-name",
						Cards: []Card{
							Card{
								ID: "test-uuid",
								Value: "Super funny new noun",
								Used: false,
							},
						},
					},
				},
				Success: true,
				Error: []errors.Error{},
				Message: "Successfully saved new card to test-name",
			},
		},
		{
			name: "request to save new card fails when POST body incorrect",
			mockService: newMockService("test-uuid"), 
			request: func() *http.Request {
				var data = []byte(`{"not_exist":"This should fail bc invalid key"}`)
				req, err := http.NewRequest("POST", "/v1/api/game/test-name/card", bytes.NewBuffer(data))
				assert.Nil(t, err)
				return req
			},
			expectedStatus: 400, 
			expectedResponse: Response{
				Result: []Game{},
				Success: false,
				Error: []errors.Error{
					errors.Error{
						Code: 1001,
						Message: "Invalid input",
					},
				},
				Message: "",
			},
		},
		{
			name: "request to save new card fails in service or repository layer",
			mockService: newMockService("error"), 
			request: func() *http.Request {
				var data = []byte(`{"value":"Super funny new noun"}`)
				req, err := http.NewRequest("POST", "/v1/api/game/test-name/card", bytes.NewBuffer(data))
				assert.Nil(t, err)
				return req
			},
			expectedStatus: 500, 
			expectedResponse: Response{
				Result: []Game{},
				Success: false,
				Error: []errors.Error{
					errors.Error{
						Code: 5000,
						Message: "Internal server error",
					},
				},
				Message: "",
			},
		},
    }

    for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var (
				response = Response{}
			)

			rr := httptest.NewRecorder()
			controller := NewGameController(tc.mockService)
			router := NewRouter(controller)
			router.ServeHTTP(rr, tc.request())

			if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
				log.Fatalln(err)
			}

			assert.Equal(t, tc.expectedStatus, rr.Code)
			assert.Equal(t, tc.expectedResponse, response)
		})
    }
}

func TestGetGameController(t *testing.T) {

	type test struct {
		name 				string
		mockService 		*mockService
		request 			func() *http.Request
		expectedStatus		int
        expectedResponse  	Response
	}
	
	tests := []test{
        {
			name: "request to fetch cards for game with one card succeeds",
			mockService: newMockService("single"), 
			request: func() *http.Request {
				req, err := http.NewRequest("GET", "/v1/api/game/single", nil)
				assert.Nil(t, err)
				return req
			},
			expectedStatus: 200, 
			expectedResponse: Response{
				Result: []Game{
					Game{
						ID: "single",
						Cards: []Card{
							Card{
								ID: "test-card",
								Value: "Funny fish",
								Used: false,
							},
						},
					},
				},
				Success: true,
				Error: []errors.Error{},
				Message: "Game single has 1 cards",
			},
		},
		{
			name: "request to fetch cards for game with no cards succeeds",
			mockService: newMockService("empty"), 
			request: func() *http.Request {
				req, err := http.NewRequest("GET", "/v1/api/game/empty", nil)
				assert.Nil(t, err)
				return req
			},
			expectedStatus: 200, 
			expectedResponse: Response{
				Result: []Game{
					Game{
						ID: "empty",
						Cards: []Card{},
					},
				},
				Success: true,
				Error: []errors.Error{},
				Message: "Game empty has 0 cards",
			},
		},
		{
			name: "request to fetch cards fails when internal error",
			mockService: newMockService("error"), 
			request: func() *http.Request {
				req, err := http.NewRequest("GET", "/v1/api/game/error", nil)
				assert.Nil(t, err)
				return req
			},
			expectedStatus: 500, 
			expectedResponse: Response{
				Result: []Game{},
				Success: false,
				Error: []errors.Error{
					errors.Error{
						Code: 5000,
						Message: "Internal server error",
					},
				},
				Message: "",
			},
		},
    }

    for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var (
				response = Response{}
			)

			rr := httptest.NewRecorder()
			controller := NewGameController(tc.mockService)
			router := NewRouter(controller)

			router.ServeHTTP(rr, tc.request())
			if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
				log.Fatalln(err)
			}

			assert.Equal(t, tc.expectedStatus, rr.Code)
			assert.Equal(t, tc.expectedResponse, response)
		})
    }
}

func newMockService(c string) *mockService {
	return &mockService{
		Case: c,
	}
}

type mockService struct {
	Case string
}

func (s *mockService) NewGame() (string, error) {
	switch s.Case {
	case "error":
		return "", fmt.Errorf("Error generating a new game")
	default:
		return s.Case, nil
	}
}

func (s *mockService) SaveCard(gameID string, card *service.CardInput) (string, error) {
	switch s.Case {
	case "error":
		return "", fmt.Errorf("Error saving new card")
	default:
		return s.Case, nil
	}
}

func (s *mockService) GetGame(gameID string) (*service.Game, error){
	game := &service.Game{}

	switch s.Case {
	case "error":
		return game, fmt.Errorf("Error fetching cards")
	case "single":
		game.ID = "single"
		game.Cards = []service.Card{
			service.Card{
				ID: "test-card",
				Value: "Funny fish",
				Used: false,
			},
		}
		return game, nil
	case "empty":
		game.ID = "empty"
		game.Cards = []service.Card{}
		return game, nil
	default:
		return game, nil
	}
}
