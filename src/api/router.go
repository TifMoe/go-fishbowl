package api

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

const (
	staticPath = "./web"
	indexPath = "index.html"
)

// NewRouter will build a new router for the routes defined below
func NewRouter(c *Controller) *mux.Router {

	r := mux.NewRouter()

	// Serve API routes
	api := r.PathPrefix("/v1/api/").Subrouter()
	api.HandleFunc("/game", c.NewGame).Methods("GET")
	api.HandleFunc("/game/{gameID}", c.FetchCards).Methods("GET")
	api.HandleFunc("/game/{gameID}/card", c.SaveNewCard).Methods("POST")

	// Serve Frontend routes
	// For requests to dynamically generated game pages, serve index.html
	r.PathPrefix("/game/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(staticPath, indexPath))
	})

	// Serve static build on root requests
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticPath)))

	// TODO - Figure out how to serve styled 404 page for unhandled paths

	return r
}
