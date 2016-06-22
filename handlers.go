package main

import (
	"log"
	"net/http"
)

// Adapter is a generic type conduit to enable middleware chaining.
type Adapter func(http.Handler) http.Handler

// AddHeader is a middleware for adding custom headers.
func AddHeader(key, value string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(key, value)
			h.ServeHTTP(w, r)
		})
	}
}

// Recover is a middleware for handling panic()s.
func Recover(logger *log.Logger) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					http.Error(w, "Unexpected error", http.StatusInternalServerError)
					logger.Println(err)
				}
			}()
			h.ServeHTTP(w, r)
		})
	}
}

// Adapt daisy-chains together multiple Adapters.
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
