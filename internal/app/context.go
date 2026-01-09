package app

import "github.com/ezh0v/pumpkin/internal/pkg/html"

type Context struct {
	AppVersion string
	html.Renderer
}
