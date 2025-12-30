package controller

import (
	"context"

	"github.com/google/uuid"
	"github.com/umisto/logium"
	"github.com/umisto/pagi"
	"github.com/umisto/profiles-svc/internal/domain/models"
	"github.com/umisto/profiles-svc/internal/domain/modules/profile"
)

type Domain interface {
	CreateProfile(ctx context.Context, userID uuid.UUID, username string) (models.Profile, error)

	FilterProfile(
		ctx context.Context,
		params profile.FilterParams,
		limit, offset uint,
	) (pagi.Page[[]models.Profile], error)

	GetProfileByID(ctx context.Context, userID uuid.UUID) (models.Profile, error)
	GetProfileByUsername(ctx context.Context, username string) (models.Profile, error)

	UpdateProfile(ctx context.Context, accountID uuid.UUID, input profile.UpdateParams) (models.Profile, error)
	UpdateProfileOfficial(ctx context.Context, accountID uuid.UUID, official bool) (models.Profile, error)
	UpdateProfileUsername(ctx context.Context, accountID uuid.UUID, username string) (models.Profile, error)
}

type Service struct {
	domain Domain
	log    logium.Logger
}

func New(log logium.Logger, profile Domain) Service {
	return Service{
		domain: profile,
		log:    log,
	}
}
