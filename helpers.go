package main

import (
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"
)

var Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	Level:     slog.LevelDebug,
	AddSource: true,
}))

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI
		trace  = string(debug.Stack())
	)

	Logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
