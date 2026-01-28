package handlers

import (
	"github.com/ezh0v/pumpkin/internal/app"
	"github.com/ezh0v/pumpkin/internal/pkg/html"
	"github.com/ezh0v/pumpkin/internal/pkg/session"
)

type Context struct {
	*html.Renderer
	*session.Manager
	*app.Instance
}

func NewContext(
	instance *app.Instance,
	renderer *html.Renderer,
	manager *session.Manager,
) *Context {
	return &Context{
		Instance: instance,
		Renderer: renderer,
		Manager:  manager,
	}
}
