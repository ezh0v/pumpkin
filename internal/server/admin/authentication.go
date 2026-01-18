package admin

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/justinas/nosurf"

	"github.com/ezh0v/pumpkin/internal/app"
	"github.com/ezh0v/pumpkin/internal/server/pkg/response"
)

type loginForm struct {
	Token    string
	Email    string
	Password string
}

func (f *loginForm) validate() error {
	return validation.ValidateStruct(f,
		validation.Field(&f.Email, validation.Required, validation.NilOrNotEmpty, is.Email),
		validation.Field(&f.Password, validation.Required, validation.NilOrNotEmpty, validation.RuneLength(10, 32)),
	)
}

func login(app *app.Context) http.HandlerFunc {
	page := renderer.NewPage("login.html")

	return func(w http.ResponseWriter, r *http.Request) {
		form := &loginForm{
			Token:    nosurf.Token(r),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		if r.Method == http.MethodGet {
			response.WithPage(w, page, "form", form)
			return
		}

		if err := form.validate(); err != nil {
			response.WithPage(w, page, "form", form, "message", err.Error())
			return
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

func logout(app *app.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
}
