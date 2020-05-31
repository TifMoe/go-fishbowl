package main

import (
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"

	"github.com/tifmoe/go-fishbowl/src/api"
	"github.com/tifmoe/go-fishbowl/src/repository"
	"github.com/tifmoe/go-fishbowl/src/service"
)

func main() {

	var (
		appPort       = api.GetEnv("PORT", "8080")
		redisHost     = api.GetEnv("REDIS_HOST", "localhost")
		redisPort     = api.GetEnv("REDIS_PORT", "6379")
		redisPassword = api.GetEnv("REDIS_PASSWORD", "")
		maxCards      = api.GetIntEnv("MAX_CARDS", 10)
	)

	// Establish redis connection
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Print("Successfully connected to the database!")

	// Instantiate repository and service layer
	repo := repository.NewRedisConnection(client)
	rand := service.NewRandomService()
	svc := service.NewGameService(repo, rand, maxCards)

	// Instantiate backend controllers and websocket router
	handlers := api.NewGameController(svc)
	pool := api.NewPool()
	go pool.Start()
	wsRouter := api.NewRouter(pool, handlers)

	r := mux.NewRouter()
	r.PathPrefix("/ws/").Handler(wsRouter)

	// Run
	if err := http.ListenAndServe(":"+appPort, r); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
