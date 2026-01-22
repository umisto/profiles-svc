package profile

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/netbill/profiles-svc/internal/core/errx"
	"github.com/netbill/profiles-svc/internal/core/models"
)

func (s Service) CreateProfile(ctx context.Context, userID uuid.UUID, username string) (profile models.Profile, err error) {
	if err = s.repo.Transaction(ctx, func(ctx context.Context) error {
		profile, err = s.repo.CreateProfile(ctx, userID, username)
		if err != nil {
			return errx.ErrorInternal.Raise(
				fmt.Errorf("creating profile for user '%s': %w", userID, err),
			)
		}

		err = s.messanger.WriteProfileCreated(ctx, profile)
		if err != nil {
			return errx.ErrorInternal.Raise(
				fmt.Errorf("sending profile created event for user '%s': %w", userID, err),
			)
		}

		return nil
	}); err != nil {
		return models.Profile{}, err
	}

	return profile, nil
}
