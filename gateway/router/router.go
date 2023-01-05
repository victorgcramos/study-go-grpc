package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Error struct {
	Code    int
	Message string
}
type PayloadHandler func([]byte) ([]byte, error)

type Route struct {
	handler PayloadHandler
	path    string
	method  string
}

type Router struct {
	// router is router
	router *mux.Router
	// protected is private, which means you need to authenticate
	protected *mux.Router
}

// NewRouter returns a new router instance
func NewRouter() *Router {
	router := mux.NewRouter()

	// TODO: setup protected credentials later
	protected := router.NewRoute().Subrouter()

	return &Router{
		router:    router,
		protected: protected,
	}
}

func respondWithBytes(w http.ResponseWriter, code int, payload []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(payload)
}

func RespondWithJSON(w http.ResponseWriter, httpcode int, payload []byte) {
	valid := json.Valid(payload)
	if valid == false {
		RespondWithError(w, 500, "invalid json encoding for payload")
	} else {
		respondWithBytes(w, httpcode, payload)
	}
}

func RespondWithError(w http.ResponseWriter, httpcode int, message string) {
	b, _ := json.Marshal(Error{Code: 1, Message: message})
	respondWithBytes(w, httpcode, b)
}

func (r *Router) AddRoute(path string, handler http.HandlerFunc, method string) {
	r.router.HandleFunc(path, handler).Methods(method)
}

func (r *Router) AddPrivateRoute(path string, handler http.HandlerFunc, method string) {
	r.protected.HandleFunc(path, handler).Methods(method)
}

func (r *Router) GetRouter() *mux.Router {
	return r.router
}
