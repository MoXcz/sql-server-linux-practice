package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./ui"))
	mux.Handle("/ui/", http.StripPrefix("/ui", fs))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /account/new", app.handleGetNewAccount)
	mux.HandleFunc("POST /account/new", app.handlePostAccount)
	mux.HandleFunc("GET /account/{accountID}", app.handleGetAccounts)
	mux.HandleFunc("GET /account", app.handleGetAccount)

	return app.recoverPanic(app.logRequest(mux))
}
