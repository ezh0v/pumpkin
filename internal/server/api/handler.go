package api

import "net/http"

func Handler() http.Handler {
	handler := http.NewServeMux()
	return http.StripPrefix("/api", handler)
}
