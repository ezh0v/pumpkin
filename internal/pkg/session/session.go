package session

import (
	"context"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/google/uuid"
)

type Manager struct {
	*scs.SessionManager
}

func NewManager(lifetime time.Duration) *Manager {
	sessionManager := scs.New()
	sessionManager.Lifetime = lifetime
	return &Manager{
		SessionManager: sessionManager,
	}
}

func (m *Manager) LoginUser(ctx context.Context, uuid uuid.UUID) error {
	if err := m.RenewToken(ctx); err != nil {
		return err
	}

	m.Put(ctx, "user", uuid.String())

	return nil
}

func (m *Manager) LogoutUser(ctx context.Context) error {
	if err := m.RenewToken(ctx); err != nil {
		return err
	}

	m.Remove(ctx, "user")

	return nil
}
