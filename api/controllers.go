package api

import (
	"fmt"
	"os"
    "net/http"
	"encoding/json"
	"path/filepath"
)

const (
	staticPath = "./web"
	indexPath = "index.html"
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

// StaticHandler is responsible for serving static react app routes
func StaticHandler(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(staticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// If dynamic route (file will not exist on server), serve index.html and let react handle routing
		http.ServeFile(w, r, filepath.Join(staticPath, indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serve static files when root requested
	fileServer := http.FileServer(http.Dir(staticPath))
	fileServer.ServeHTTP(w, r)
}