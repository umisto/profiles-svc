package profile

import (
	"context"

	"github.com/google/uuid"
	"github.com/netbill/profiles-svc/internal/core/models"
)

func (m *Module) GetProfileByAccountID(ctx context.Context, userID uuid.UUID) (models.Profile, error) {
	return m.repo.GetProfileByAccountID(ctx, userID)
}

func (m *Module) GetProfileByUsername(ctx context.Context, username string) (models.Profile, error) {
	return m.repo.GetProfileByUsername(ctx, username)
}
