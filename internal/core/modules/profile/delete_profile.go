package profile

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/netbill/profiles-svc/internal/core/errx"
)

func (s Service) DeleteProfile(ctx context.Context, userID uuid.UUID) error {
	return s.repo.Transaction(ctx, func(ctx context.Context) error {
		err := s.repo.DeleteProfile(ctx, userID)
		if err != nil {
			return errx.ErrorInternal.Raise(
				fmt.Errorf("failed to delete profile: %w", err),
			)
		}

		err = s.messanger.WriteProfileDeleted(ctx, userID)
		if err != nil {
			return errx.ErrorInternal.Raise(
				fmt.Errorf("failed to publish profile deleted event: %w", err),
			)
		}

		return nil
	})
}
