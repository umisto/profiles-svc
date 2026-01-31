package profile

import (
	"context"

	"github.com/google/uuid"
	"github.com/netbill/profiles-svc/internal/core/models"
)

func (s Service) UpdateProfileUsername(ctx context.Context, accountID uuid.UUID, username string) (profile models.Profile, err error) {
	if err = s.repo.Transaction(ctx, func(ctx context.Context) error {
		profile, err = s.repo.UpdateProfileUsername(ctx, accountID, username)
		if err != nil {
			return err
		}

		err = s.messanger.WriteProfileUpdated(ctx, profile)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return models.Profile{}, err
	}

	return profile, nil
}
