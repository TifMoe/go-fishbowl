package api

import (
    "net/http"
	"encoding/json"
)

// Ping is Handler for Pinging Server
func Ping(w http.ResponseWriter, r *http.Request) {
	res := &Response{
		Message: "I'm aaaalivvveee!!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
