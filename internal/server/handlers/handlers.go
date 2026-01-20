package handlers

import (
	"html/template"

	"github.com/alexedwards/scs/v2"

	"github.com/ezh0v/pumpkin/internal/app"
	"github.com/ezh0v/pumpkin/internal/pkg/html"
)

type Context struct {
	*app.Context
	*html.Renderer
	*scs.SessionManager
}

func NewContext(app *app.Context, views *template.Template) *Context {
	return &Context{
		Context:        app,
		Renderer:       html.NewRenderer(views),
		SessionManager: scs.New(),
	}
}
