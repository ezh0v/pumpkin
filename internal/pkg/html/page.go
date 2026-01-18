package html

import "io"

type Page struct {
	name     string
	renderer *Renderer
}

func (p *Page) Render(w io.Writer, args ...any) error {
	return p.renderer.RenderTemplate(w, p.name, args...)
}
