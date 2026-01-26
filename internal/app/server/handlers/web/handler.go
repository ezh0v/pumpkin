package web

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/ezh0v/pumpkin/internal/app"
	"github.com/ezh0v/pumpkin/internal/app/server/handlers"
	"github.com/ezh0v/pumpkin/internal/app/server/middlewares"
)

//go:embed views
var viewsFS embed.FS

func Handler(app *app.Instance) http.Handler {
	c := handlers.NewContext(app,
		template.Must(template.ParseFS(viewsFS, "views/**/*.html")),
	)

	handler := http.NewServeMux()
	handler.HandleFunc("/", home(c))

	return middlewares.With(handler,
		c.LoadAndSave,
		middlewares.CSRF(),
	)
}
