package profile

import (
	"context"

	"github.com/google/uuid"
)

func (m *Module) DeleteProfile(ctx context.Context, userID uuid.UUID) error {
	return m.repo.Transaction(ctx, func(ctx context.Context) error {
		err := m.repo.DeleteProfile(ctx, userID)
		if err != nil {
			return err
		}

		err = m.messanger.WriteProfileDeleted(ctx, userID)
		if err != nil {
			return err
		}

		return nil
	})
}
