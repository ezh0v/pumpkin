package admin

import (
	"net/http"

	"github.com/ezh0v/pumpkin/internal/app"
)

func Handler(app *app.Context) http.Handler {
	handler := http.NewServeMux()
	return http.StripPrefix("/admin", handler)
}
