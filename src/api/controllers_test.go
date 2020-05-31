package api

// TODO: Update test cases below to work with WebSockets

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"

// 	"github.com/tifmoe/go-fishbowl/src/errors"
// 	"github.com/tifmoe/go-fishbowl/src/service"
// )

// func TestNewGameController(t *testing.T) {

// 	type test struct {
// 		name             string
// 		mockService      *mockService
// 		postData         json.RawMessage
// 		expectedStatus   int
// 		expectedResponse Response
// 	}

// 	tests := []test{
// 		{
// 			name:           "request to create new game succeeds",
// 			mockService:    newMockService("successful-game"),
// 			postData:       json.RawMessage(`{"team1": "boo", "team2": "hoo"}`),
// 			expectedStatus: 201,
// 			expectedResponse: Response{
// 				Result: []Game{
// 					Game{
// 						ID: "successful-game",
// 					},
// 				},
// 				Success: true,
// 				Error:   []errors.Error{},
// 				Message: "successful-game",
// 			},
// 		},
// 		{
// 			name:           "request to create new game fails on internal service layer error",
// 			mockService:    newMockService("error"),
// 			postData:       json.RawMessage(`{"team1": "boo", "team2": "hoo"}`),
// 			expectedStatus: 500,
// 			expectedResponse: Response{
// 				Result:  []Game{},
// 				Success: false,
// 				Error: []errors.Error{
// 					errors.Error{
// 						Code:    1000,
// 						Message: "Failed to instantiate new game session, please try again later",
// 					},
// 				},
// 				Message: "",
// 			},
// 		},
// 		{
// 			name:           "request to create new game fails when team names are not provided",
// 			mockService:    newMockService("error"),
// 			postData:       json.RawMessage(`{"invalid": "misguided input"}`),
// 			expectedStatus: 500,
// 			expectedResponse: Response{
// 				Result:  []Game{},
// 				Success: false,
// 				Error: []errors.Error{
// 					errors.Error{
// 						Code:    1000,
// 						Message: "Failed to instantiate new game session, please try again later",
// 					},
// 				},
// 				Message: "",
// 			},
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			var (
// 				response = Response{}
// 			)

// 			req, err := http.NewRequest("POST", "/v1/api/game", bytes.NewReader(tc.postData))
// 			assert.Nil(t, err)

// 			rr := httptest.NewRecorder()
// 			controller := NewGameController(tc.mockService)
// 			handler := http.HandlerFunc(controller.NewGame)
// 			handler.ServeHTTP(rr, req)

// 			if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
// 				log.Fatalln(err)
// 			}

// 			assert.Equal(t, tc.expectedStatus, rr.Code)
// 			assert.Equal(t, tc.expectedResponse, response)
// 		})
// 	}
// }

// func TestNewCardController(t *testing.T) {

// 	type test struct {
// 		name             string
// 		mockService      *mockService
// 		request          func() *http.Request
// 		expectedStatus   int
// 		expectedResponse Response
// 	}

// 	tests := []test{
// 		{
// 			name:        "request to save new card succeeds",
// 			mockService: newMockService("test-uuid"),
// 			request: func() *http.Request {
// 				var data = []byte(`{"value":"Super funny new noun"}`)
// 				req, err := http.NewRequest("POST", "/v1/api/game/test-name/card", bytes.NewBuffer(data))
// 				assert.Nil(t, err)
// 				return req
// 			},
// 			expectedStatus: 201,
// 			expectedResponse: Response{
// 				Result: []Game{
// 					Game{
// 						ID: "test-name",
// 						Cards: []Card{
// 							Card{
// 								ID:    "test-uuid",
// 								Value: "Super funny new noun",
// 								Used:  false,
// 							},
// 						},
// 					},
// 				},
// 				Success: true,
// 				Error:   []errors.Error{},
// 				Message: "Successfully saved new card to test-name",
// 			},
// 		},
// 		{
// 			name:        "request to save new card fails when POST body incorrect",
// 			mockService: newMockService("test-uuid"),
// 			request: func() *http.Request {
// 				var data = []byte(`{"not_exist":"This should fail bc invalid key"}`)
// 				req, err := http.NewRequest("POST", "/v1/api/game/test-name/card", bytes.NewBuffer(data))
// 				assert.Nil(t, err)
// 				return req
// 			},
// 			expectedStatus: 400,
// 			expectedResponse: Response{
// 				Result:  []Game{},
// 				Success: false,
// 				Error: []errors.Error{
// 					errors.Error{
// 						Code:    1001,
// 						Message: "Invalid input",
// 					},
// 				},
// 				Message: "",
// 			},
// 		},
// 		{
// 			name:        "request to save new card fails in service or repository layer",
// 			mockService: newMockService("error"),
// 			request: func() *http.Request {
// 				var data = []byte(`{"value":"Super funny new noun"}`)
// 				req, err := http.NewRequest("POST", "/v1/api/game/test-name/card", bytes.NewBuffer(data))
// 				assert.Nil(t, err)
// 				return req
// 			},
// 			expectedStatus: 500,
// 			expectedResponse: Response{
// 				Result:  []Game{},
// 				Success: false,
// 				Error: []errors.Error{
// 					errors.Error{
// 						Code:    5000,
// 						Message: "Internal server error",
// 					},
// 				},
// 				Message: "",
// 			},
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			var (
// 				response = Response{}
// 			)

// 			rr := httptest.NewRecorder()
// 			controller := NewGameController(tc.mockService)
// 			router := NewRouter(controller)
// 			router.ServeHTTP(rr, tc.request())

// 			if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
// 				log.Fatalln(err)
// 			}

// 			assert.Equal(t, tc.expectedStatus, rr.Code)
// 			assert.Equal(t, tc.expectedResponse, response)
// 		})
// 	}
// }

// func TestGetGameController(t *testing.T) {

// 	type test struct {
// 		name             string
// 		mockService      *mockService
// 		request          func() *http.Request
// 		expectedStatus   int
// 		expectedResponse Response
// 	}

// 	tests := []test{
// 		{
// 			name:        "request to fetch cards for game with one card succeeds",
// 			mockService: newMockService("single"),
// 			request: func() *http.Request {
// 				req, err := http.NewRequest("GET", "/v1/api/game/single", nil)
// 				assert.Nil(t, err)
// 				return req
// 			},
// 			expectedStatus: 200,
// 			expectedResponse: Response{
// 				Result: []Game{
// 					Game{
// 						ID: "single",
// 						Cards: []Card{
// 							Card{
// 								ID:    "test-card",
// 								Value: "Pants king",
// 								Used:  false,
// 							},
// 						},
// 					},
// 				},
// 				Success: true,
// 				Error:   []errors.Error{},
// 				Message: "Game single has 1 cards",
// 			},
// 		},
// 		{
// 			name:        "request to fetch cards for game with no cards succeeds",
// 			mockService: newMockService("empty"),
// 			request: func() *http.Request {
// 				req, err := http.NewRequest("GET", "/v1/api/game/empty", nil)
// 				assert.Nil(t, err)
// 				return req
// 			},
// 			expectedStatus: 200,
// 			expectedResponse: Response{
// 				Result: []Game{
// 					Game{
// 						ID:    "empty",
// 						Cards: []Card{},
// 					},
// 				},
// 				Success: true,
// 				Error:   []errors.Error{},
// 				Message: "Game empty has 0 cards",
// 			},
// 		},
// 		{
// 			name:        "request to fetch cards fails when internal error",
// 			mockService: newMockService("error"),
// 			request: func() *http.Request {
// 				req, err := http.NewRequest("GET", "/v1/api/game/error", nil)
// 				assert.Nil(t, err)
// 				return req
// 			},
// 			expectedStatus: 500,
// 			expectedResponse: Response{
// 				Result:  []Game{},
// 				Success: false,
// 				Error: []errors.Error{
// 					errors.Error{
// 						Code:    5000,
// 						Message: "Internal server error",
// 					},
// 				},
// 				Message: "",
// 			},
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			var (
// 				response = Response{}
// 			)

// 			rr := httptest.NewRecorder()
// 			controller := NewGameController(tc.mockService)
// 			router := NewRouter(controller)

// 			router.ServeHTTP(rr, tc.request())
// 			if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
// 				log.Fatalln(err)
// 			}

// 			assert.Equal(t, tc.expectedStatus, rr.Code)
// 			assert.Equal(t, tc.expectedResponse, response)
// 		})
// 	}
// }

// func TestGetRandomCardController(t *testing.T) {

// 	type test struct {
// 		name             string
// 		mockService      *mockService
// 		expectedStatus   int
// 		expectedResponse Response
// 	}

// 	tests := []test{
// 		{
// 			name:           "request to fetch random card for game succeeds",
// 			mockService:    newMockService("single"),
// 			expectedStatus: 200,
// 			expectedResponse: Response{
// 				Result: []Game{
// 					Game{
// 						ID: "single",
// 						Cards: []Card{
// 							Card{
// 								ID:    "random-card",
// 								Value: "Trump's carrot fingers",
// 								Used:  false,
// 							},
// 						},
// 					},
// 				},
// 				Success: true,
// 				Error:   []errors.Error{},
// 				Message: "",
// 			},
// 		},
// 		{
// 			name:           "request to fetch random card when no unused cards left is successful",
// 			mockService:    newMockService("empty"),
// 			expectedStatus: 200,
// 			expectedResponse: Response{
// 				Result: []Game{
// 					Game{
// 						ID:    "empty",
// 						Cards: []Card{},
// 					},
// 				},
// 				Success: true,
// 				Error:   []errors.Error{},
// 				Message: "",
// 			},
// 		},
// 		{
// 			name:           "request to fetch random card fails when internal error",
// 			mockService:    newMockService("error"),
// 			expectedStatus: 500,
// 			expectedResponse: Response{
// 				Result:  []Game{},
// 				Success: false,
// 				Error: []errors.Error{
// 					errors.Error{
// 						Code:    5000,
// 						Message: "Internal server error",
// 					},
// 				},
// 				Message: "",
// 			},
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			var (
// 				response = Response{}
// 			)

// 			rr := httptest.NewRecorder()
// 			controller := NewGameController(tc.mockService)
// 			router := NewRouter(controller)

// 			req, err := http.NewRequest("GET", "/v1/api/game/game-name/card/random", nil)
// 			assert.Nil(t, err)
// 			router.ServeHTTP(rr, req)
// 			if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
// 				log.Fatalln(err)
// 			}

// 			assert.Equal(t, tc.expectedStatus, rr.Code)
// 			assert.Equal(t, tc.expectedResponse, response)
// 		})
// 	}
// }

// func newMockService(c string) *mockService {
// 	return &mockService{
// 		Case: c,
// 	}
// }

// type mockService struct {
// 	Case string
// }

// func (s *mockService) NewGame(input *service.TeamInput) (string, error) {
// 	switch s.Case {
// 	case "error":
// 		return "", fmt.Errorf("Error generating a new game")
// 	default:
// 		return s.Case, nil
// 	}
// }

// func (s *mockService) SaveCard(gameID string, card *service.CardInput) (string, error) {
// 	switch s.Case {
// 	case "error":
// 		return "", fmt.Errorf("Error saving new card")
// 	default:
// 		return s.Case, nil
// 	}
// }

// func (s *mockService) GetGame(gameID string) (*service.Game, error) {
// 	game := &service.Game{}

// 	switch s.Case {
// 	case "error":
// 		return game, fmt.Errorf("Error fetching cards")
// 	case "single":
// 		game.ID = "single"
// 		game.Cards = []service.Card{
// 			service.Card{
// 				ID:    "test-card",
// 				Value: "Pants king",
// 				Used:  false,
// 			},
// 		}
// 		return game, nil
// 	case "empty":
// 		game.ID = "empty"
// 		game.Cards = []service.Card{}
// 		return game, nil
// 	default:
// 		return game, nil
// 	}
// }

// func (s *mockService) GetRandomCard(gameID string) (card *service.Card, err error) {
// 	switch s.Case {
// 	case "error":
// 		err = fmt.Errorf("Error fetching random card")
// 		return
// 	case "single":
// 		card = &service.Card{
// 			ID:    "random-card",
// 			Value: "Trump's carrot fingers",
// 			Used:  false,
// 		}
// 		return
// 	}
// 	return
// }

// func (s *mockService) MarkCardUsed(gameID, cardID string) (*service.Game, error) {
// 	return &service.Game{}, nil
// }

// func (s *mockService) ResetGame(gameID string) (*service.Game, error) {
// 	return &service.Game{}, nil
// }

// func (s *mockService) DeleteCards(gameID string) error {
// 	return nil
// }

// func (s *mockService) UpdateGame(gameID string, input *service.GameInput) (*service.Game, error) {
// 	var game *service.Game
// 	return game, nil
// }
