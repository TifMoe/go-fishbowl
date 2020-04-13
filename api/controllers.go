package api

import (
	"fmt"
    "net/http"
	"encoding/json"
)

// TODO Instantiate controller interface with randomWord service to make testing easier

// RandomWords is controller for generating new pair of random words
func RandomWords(w http.ResponseWriter, r *http.Request) {
	nameSpace, err := GetRandomWords()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	res := &Response{
		Message: nameSpace,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
