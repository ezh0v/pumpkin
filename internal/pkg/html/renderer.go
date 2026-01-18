package html

import (
	"errors"
	"fmt"
	"html/template"
	"io"
)

type Renderer struct {
	templates *template.Template
}

func NewRenderer(templates *template.Template) *Renderer {
	return &Renderer{
		templates: templates,
	}
}

func (r *Renderer) NewPage(name string) *Page {
	return &Page{
		renderer: r,
		name:     name,
	}
}

func (r *Renderer) RenderTemplate(w io.Writer, name string, args ...any) error {
	data := make(map[string]any)

	if len(args)%2 != 0 {
		return errors.New("args must be even number of elements: key, value...")
	}

	for i := 0; i < len(args); i += 2 {
		key, ok := args[i].(string)
		if !ok {
			return fmt.Errorf("argument %d is not a string key", i)
		}

		data[key] = args[i+1]
	}

	return r.templates.ExecuteTemplate(w, name, data)
}
