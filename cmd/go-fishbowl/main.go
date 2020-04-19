package main

import (
    "net/http"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"gopkg.in/go-playground/validator.v9"

	"github.com/tifmoe/go-fishbowl/src/api"
	"github.com/tifmoe/go-fishbowl/src/service"
	"github.com/tifmoe/go-fishbowl/src/repository"
)

func main() {
	port := api.GetEnv("PORT", "8080")
	password := api.GetEnv("REDIS_PASSWORD", "")

	// Establish redis connection
	client := redis.NewClient(&redis.Options{
		Addr: "db:6379",
		Password: password,
		DB: 0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Print("Successfully connected to the database!")

	// Instantiate repository and service layer
	v := validator.New()
	repo := repository.NewRedisConnection(client)
	svc := service.NewGameService(repo, v)

	// Instantiate controllers and router
	handlers := api.NewGameController(svc)
	router := api.NewRouter(handlers)

	// Run
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}