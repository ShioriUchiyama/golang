package main

import (
	"chat-websockets/internal/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/chat/", http.HandlerFunc(handlers.Chat))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))
	return mux
}
