package contracts

import (
	"time"

	"github.com/google/uuid"
)

const ProfileUpdatedEvent = "profile.updated"

type AccountProfileUpdatedPayload struct {
	AccountID   uuid.UUID `json:"account_id"`
	Pseudonym   *string   `json:"pseudonym,omitempty"`
	Description *string   `json:"description,omitempty"`

	UpdatedAt time.Time `json:"updated_at"`
}

const ProfileOfficialUpdatedEvent = "profile.official.updated"

type AccountProfileOfficialUpdatedPayload struct {
	AccountID uuid.UUID `json:"account_id"`
	Official  bool      `json:"official"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
