package profile

import (
	"context"

	"github.com/google/uuid"
	"github.com/netbill/profiles-svc/internal/core/models"
)

func (s Service) UpdateProfileOfficial(ctx context.Context, accountID uuid.UUID, official bool) (profile models.Profile, err error) {
	if err = s.repo.Transaction(ctx, func(ctx context.Context) error {
		profile, err = s.repo.UpdateProfileOfficial(ctx, accountID, official)
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
