package main

import (
    "net/http"
	"log"

	"github.com/go-redis/redis"
	"gopkg.in/go-playground/validator.v9"

	"github.com/tifmoe/go-fishbowl/src/api"
	"github.com/tifmoe/go-fishbowl/src/service"
	"github.com/tifmoe/go-fishbowl/src/repository"
)

func main() {

	var (
		appPort 	= api.GetEnv("PORT", "8080")
		redisHost     	= api.GetEnv("REDIS_HOST", "localhost")
		redisPort     	= api.GetEnv("REDIS_PORT", "6379")
		redisPassword 	= api.GetEnv("REDIS_PASSWORD", "")
		maxCards 	= api.GetIntEnv("MAX_CARDS", 10)
	)

	// Establish redis connection
	client := redis.NewClient(&redis.Options{
		Addr: redisHost + ":" + redisPort,
		Password: redisPassword,
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
	svc := service.NewGameService(repo, v, maxCards)

	// Instantiate controllers and router
	handlers := api.NewGameController(svc)
	router := api.NewRouter(handlers)

	// Run
	if err := http.ListenAndServe(":" + appPort, router); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}