package contracts

import "github.com/umisto/profiles-svc/internal/domain/models"

const ProfilesTopicV1 = "profiles.v1"

const ProfileUpdatedEvent = "profile.updated"

type ProfileUpdatedPayload struct {
	Profile models.Profile `json:"profile"`
}
