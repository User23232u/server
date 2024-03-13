package routes

import (
	"net/http"

	"project-websocket/websockets"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Aceptar todas las solicitudes
	},
}

// @summary Create a new router
// @description Create a new router with route handlers.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ws", websockets.WSHandler)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./docs")))
	return router
}