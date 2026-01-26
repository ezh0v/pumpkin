package middlewares

import (
	"net/http"

	"github.com/justinas/nosurf"
)

type Middleware func(http.Handler) http.Handler

func With(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}

func CSRF() Middleware {
	return func(next http.Handler) http.Handler {
		// TODO: configure
		return nosurf.New(next)
	}
}
