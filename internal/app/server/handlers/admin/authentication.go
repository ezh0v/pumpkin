package admin

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/justinas/nosurf"

	"github.com/ezh0v/pumpkin/internal/app/server/handlers"
	"github.com/ezh0v/pumpkin/internal/app/server/response"
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

func login(c *handlers.Context) http.HandlerFunc {
	page := c.NewPage("login.html")

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
			response.WithPage(w, page, "form", form)
			return
		}

		ctx := r.Context()

		if err := c.RenewToken(ctx); err != nil {
			response.WithPage(w, page, "form", form)
			return
		}

		c.Put(ctx, "userUUID", "")

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

func logout(c *handlers.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}

		ctx := r.Context()

		if err := c.RenewToken(ctx); err != nil {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}

		c.Remove(ctx, "userUUID")

		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
}
