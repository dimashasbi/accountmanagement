package adapter

import (
	"AccountManagement/engine"

	"github.com/gorilla/mux"

	"net/http"
)

// Handler using for make a Route
type (
	// Handler structure for Application Start Server
	Handler struct {
		Router   *mux.Router
		muxUsers *users
	}
)

// InitializeServer Application
func (a *Handler) InitializeServer(f engine.EnginesFactory) {
	// add Engine
	a.muxUsers = &users{f.NewUsersEngine()}
	a.Router = mux.NewRouter()
	a.SetURL()
}

// SetURL for reloading
func (a *Handler) SetURL() {

}

// GET wraps the router for GET method
func (a *Handler) GET(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// POST wraps the router for POST method
func (a *Handler) POST(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Run the app on it's router
func (a *Handler) Run(port string) {
	http.ListenAndServe(port, a.Router)
}
