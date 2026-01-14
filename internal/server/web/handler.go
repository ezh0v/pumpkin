package web

import (
	"embed"
	"html/template"
	"net/http"
)

var (
	//go:embed views
	viewsFS embed.FS

	templates *template.Template
)

func init() {
	templates = template.Must(template.ParseFS(viewsFS, "views/**/*.html"))
}

func Handler() http.Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("GET /", home())
	return handler
}
