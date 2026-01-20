package handlers

import (
	"html/template"

	"github.com/alexedwards/scs/v2"

	"github.com/ezh0v/pumpkin/internal/app"
	"github.com/ezh0v/pumpkin/internal/pkg/html"
)

type Context struct {
	*html.Renderer
	*scs.SessionManager
	*app.Context
}

func NewContext(app *app.Context, templates *template.Template) *Context {
	return &Context{
		Context:        app,
		Renderer:       html.NewRenderer(templates),
		SessionManager: scs.New(),
	}
}
