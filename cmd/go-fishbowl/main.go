package main

import (
    "net/http"

	"github.com/tifmoe/go-fishbowl/api"
)

func main() {

	port := api.GetEnv("PORT", "8080")

	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/v1/api/random/name", api.RandomWords)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}