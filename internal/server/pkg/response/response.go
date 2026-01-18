package response

import (
	"net/http"

	"github.com/ezh0v/pumpkin/internal/pkg/html"
)

func WithPage(w http.ResponseWriter, page *html.Page, args ...any) {
	if err := page.Render(w, args...); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
