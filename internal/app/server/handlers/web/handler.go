package web

import (
	"embed"
	"html/template"
	"net/http"
	"time"

	"github.com/ezh0v/pumpkin/internal/app"
	"github.com/ezh0v/pumpkin/internal/app/server/handlers"
	"github.com/ezh0v/pumpkin/internal/app/server/middlewares"
	"github.com/ezh0v/pumpkin/internal/pkg/html"
	"github.com/ezh0v/pumpkin/internal/pkg/session"
)

//go:embed views
var viewsFS embed.FS

func Handler(app *app.Instance) http.Handler {
	templates := template.Must(template.ParseFS(viewsFS, "views/**/*.html"))

	renderer := html.NewRenderer(templates,
		html.WithGlobalValue("appVersion", app.Version),
	)
	manager := session.NewManager(48 * time.Hour)

	context := handlers.NewContext(app, renderer, manager)

	handler := http.NewServeMux()
	handler.HandleFunc("/", home(context))

	return middlewares.With(handler,
		manager.LoadAndSave,
		middlewares.CSRF(),
	)
}
