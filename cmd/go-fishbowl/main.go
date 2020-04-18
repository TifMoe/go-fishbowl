package main

import (
    "net/http"

	"github.com/tifmoe/go-fishbowl/api"
)


func main() {
	port := api.GetEnv("PORT", "8080")

	// Serve API routes
	http.HandleFunc("/v1/api/random/name", api.RandomWords)

	// Serve frontend routes
	http.HandleFunc("/", api.StaticHandler)

	// Run
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}