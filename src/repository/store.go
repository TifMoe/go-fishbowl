package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const (
	ttl = 120 * time.Minute // Two hour time to live for games (after last update)
)

// NewRedisConnection will instantiate a new connection to the redis db
func NewRedisConnection(c *redis.Client) *conn {
	return &conn{
		Client: c,
	}
}

// Repository is interface with methods to interact with redis db	
type Repository interface {
    SaveNewGame(gameID, team1, team2 string) error
	UpdateGame(game *Game) error
	GetGame(gameID string) (game *Game, err error)
}

// conn is a struct holding the redis client
type conn struct {
    Client *redis.Client
}

// SaveNewGame is used to initialize a new game
func (c conn) SaveNewGame(gameID, team1, team2 string) error {
	exists, _ := c.GetGame(gameID)

	if exists != nil {
		err := fmt.Errorf("failed to save: game %s already exists", gameID)
		return err
	}

	game, err := json.Marshal(Game{
		ID: gameID,
		Teams: Teams{
			Team1: Team{
				Name: team1,
			},
			Team2: Team{
				Name: team2,
			},
		},
	})

	if err != nil {
		fmt.Printf("Error marshalling new game %s: %v\n", gameID, err)
		return err
	}
    err = c.Client.Set(gameID, game, ttl).Err()
    if err != nil {
		fmt.Printf("Error saving game %s: %v\n", gameID, err)
        return err
	}
	return nil
}

func (c conn) GetGame(gameID string) (game *Game, err error) {
    data, err := c.Client.Get(gameID).Result()
    if err != nil {
		fmt.Printf("Game %s does not exist: %v\n", gameID, err)
		return game, err
    }

	err = json.Unmarshal([]byte(data), &game)
    if err != nil {
		fmt.Printf("Error marshalling game %s: %v\n", gameID, err)
		return game, err
	}
	return game, nil
}

func (c conn) UpdateGame(game *Game) error {
	gameID := game.ID
	updatedGame, err := json.Marshal(game)
    if err != nil {
		fmt.Printf("Error marshalling updated game %s: %v\n", gameID, err)
		return err
	}

    err = c.Client.Set(gameID, updatedGame, ttl).Err()
    if err != nil {
		fmt.Printf("Error saving game %v: %v\n", gameID, err)
		return err
    }
	return nil
}
