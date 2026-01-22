package html

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"maps"
)

type Renderer struct {
	globalValues map[string]any
	templates    *template.Template
}

func NewRenderer(templates *template.Template, opts ...Option) *Renderer {
	options := &options{
		globalValues: make(map[string]any),
	}

	for _, opt := range opts {
		opt(options)
	}

	return &Renderer{
		globalValues: options.globalValues,
		templates:    templates,
	}
}

func (r *Renderer) NewPage(name string) *Page {
	return &Page{
		renderer: r,
		name:     name,
	}
}

func (r *Renderer) RenderTemplate(w io.Writer, name string, args ...any) error {
	if len(args)%2 != 0 {
		return errors.New("args must be even number of elements: key, value...")
	}

	data := maps.Clone(r.globalValues)

	for i := 0; i < len(args); i += 2 {
		key, ok := args[i].(string)
		if !ok {
			return fmt.Errorf("argument %d is not a string key", i)
		}

		data[key] = args[i+1]
	}

	return r.templates.ExecuteTemplate(w, name, data)
}
