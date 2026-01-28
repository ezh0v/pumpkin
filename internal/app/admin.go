package app

import (
	"context"

	"github.com/ezh0v/pumpkin/internal/app/models"
)

type Admin struct {
	*Instance
}

func (a *Admin) Authenticate(ctx context.Context, email, password string) (*models.User, error) {
	return &models.User{}, nil
}
