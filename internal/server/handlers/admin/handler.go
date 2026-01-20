package admin

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/ezh0v/pumpkin/internal/app"
	"github.com/ezh0v/pumpkin/internal/server/handlers"
	"github.com/ezh0v/pumpkin/internal/server/middlewares"
)

//go:embed views
var viewsFS embed.FS

func Handler(app *app.Instance) http.Handler {
	c := handlers.NewContext(app,
		template.Must(template.ParseFS(viewsFS, "views/**/*.html")),
	)

	handler := http.NewServeMux()
	handler.HandleFunc("/", home(c))
	handler.HandleFunc("/login", login(c))
	handler.HandleFunc("/logout", logout(c))

	return http.StripPrefix("/admin", middlewares.With(handler,
		c.LoadAndSave,
		middlewares.CSRF(),
	))
}
