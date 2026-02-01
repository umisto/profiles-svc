package profile

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/netbill/profiles-svc/internal/core/errx"
	"github.com/netbill/profiles-svc/internal/core/models"
)

func (m *Module) OpenProfileUpdateSession(
	ctx context.Context,
	accountID uuid.UUID,
) (models.UpdateProfileMedia, models.Profile, error) {
	profile, err := m.GetProfileByAccountID(ctx, accountID)
	if err != nil {
		return models.UpdateProfileMedia{}, models.Profile{}, err
	}

	uploadSessionID := uuid.New()
	links, err := m.bucket.GetPreloadLinkForProfileMedia(
		ctx,
		accountID,
		uploadSessionID,
	)
	if err != nil {
		return models.UpdateProfileMedia{}, models.Profile{}, err
	}

	uploadToken, err := m.token.NewUploadProfileMediaToken(accountID, uploadSessionID)
	if err != nil {
		return models.UpdateProfileMedia{}, models.Profile{}, err
	}

	return models.UpdateProfileMedia{
		Links:           links,
		UploadSessionID: uploadSessionID,
		UploadToken:     uploadToken,
	}, profile, nil
}

type UpdateParams struct {
	Pseudonym   *string
	Description *string

	Media UpdateMediaParams
}

type UpdateMediaParams struct {
	UploadSessionID uuid.UUID

	DeleteAvatar bool
	avatarKey    *string
}

func (p UpdateParams) GetUpdatedAvatar() *string {
	if p.Media.DeleteAvatar {
		return nil
	}

	return p.Media.avatarKey
}

func (m *Module) UpdateProfile(
	ctx context.Context,
	accountID uuid.UUID,
	params UpdateParams,
) (profile models.Profile, err error) {
	profile, err = m.GetProfileByAccountID(ctx, accountID)
	if err != nil {
		return models.Profile{}, err
	}

	params.Media.avatarKey = profile.Avatar
	switch params.Media.DeleteAvatar {
	case true:
		if err = m.bucket.DeleteProfileAvatar(
			ctx,
			accountID,
		); err != nil {
			return models.Profile{}, err
		}

		params.Media.avatarKey = nil
	case false:
		avatar, err := m.bucket.AcceptUpdateProfileMedia(
			ctx,
			accountID,
			params.Media.UploadSessionID,
		)
		switch {
		case errors.Is(err, errx.ErrorNoContentUploaded):
			// No new avatar uploaded, keep the existing one
		case err != nil:
			return models.Profile{}, err
		default:
			params.Media.avatarKey = &avatar
		}
	}

	err = m.bucket.CleanProfileMediaSession(
		ctx,
		accountID,
		params.Media.UploadSessionID,
	)
	if err != nil {
		return models.Profile{}, err
	}

	if err = m.repo.Transaction(ctx, func(ctx context.Context) error {
		profile, err = m.repo.UpdateProfile(ctx, accountID, params)
		if err != nil {
			return err
		}

		err = m.messanger.WriteProfileUpdated(ctx, profile)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return models.Profile{}, err
	}

	return profile, nil
}

func (m *Module) DeleteUploadProfileAvatarInSession(
	ctx context.Context,
	accountID, sessionID uuid.UUID,
) error {
	err := m.bucket.CancelUpdateProfileAvatar(ctx, accountID, sessionID)
	if err != nil {
		return err
	}

	return nil
}
