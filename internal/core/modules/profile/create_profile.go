package profile

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/netbill/profiles-svc/internal/core/errx"
	"github.com/netbill/profiles-svc/internal/core/models"
)

func (s Service) CreateProfile(ctx context.Context, accountID uuid.UUID, username string) (models.Profile, error) {
	profile, err := s.repo.GetProfileByAccountID(ctx, accountID)
	if err != nil {
		return models.Profile{}, errx.ErrorInternal.Raise(
			fmt.Errorf("checking existing profile for user '%s': %w", accountID, err),
		)
	}
	if !profile.IsNil() {
		return profile, nil
	}

	if err = s.repo.Transaction(ctx, func(ctx context.Context) error {
		profile, err = s.repo.InsertProfile(ctx, accountID, username)
		if err != nil {
			return errx.ErrorInternal.Raise(
				fmt.Errorf("creating profile for user '%s': %w", accountID, err),
			)
		}

		err = s.messanger.WriteProfileCreated(ctx, profile)
		if err != nil {
			return errx.ErrorInternal.Raise(
				fmt.Errorf("sending profile created event for user '%s': %w", accountID, err),
			)
		}

		return nil
	}); err != nil {
		return models.Profile{}, err
	}

	return profile, nil
}
