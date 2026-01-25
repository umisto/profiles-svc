package pgdb

import (
	"github.com/netbill/profiles-svc/internal/core/models"
)

func (p *Profile) ToModel() models.Profile {
	profile := models.Profile{
		AccountID: p.AccountID,
		Username:  p.Username,
		Official:  p.Official,

		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
	if p.Pseudonym.Valid {
		profile.Pseudonym = &p.Pseudonym.String
	}
	if p.Description.Valid {
		profile.Description = &p.Description.String
	}
	if p.AvatarURL.Valid {
		profile.AvatarURL = &p.AvatarURL.String
	}
	return profile
}
