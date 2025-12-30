package pgdb

import (
	"github.com/umisto/profiles-svc/internal/domain/models"
)

func (p Profile) ToModel() models.Profile {
	profile := models.Profile{
		AccountID:   p.AccountID,
		Username:    p.Username,
		Official:    p.Official,
		Pseudonym:   p.Pseudonym,
		Description: p.Description,
		Avatar:      p.Avatar,

		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
	return profile
}
