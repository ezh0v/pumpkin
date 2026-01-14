package admin

import (
	"net/http"
)

func Handler() http.Handler {
	handler := http.NewServeMux()
	return http.StripPrefix("/admin", handler)
}
