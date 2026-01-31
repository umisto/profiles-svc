package tokenmanager

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/netbill/restkit/tokens"
)

type Manager struct {
	uploadSK string

	profileMediaUploadTTL time.Duration
}

const (
	ProfilesActor   = "profiles-svc"
	ProfileResource = "profile"
)

func New(uploadSK string, profileMediaUploadTTL time.Duration) Manager {
	return Manager{
		uploadSK:              uploadSK,
		profileMediaUploadTTL: profileMediaUploadTTL,
	}
}

func (m Manager) NewUploadProfileMediaToken(
	OwnerAccountID uuid.UUID,
	UploadSessionID uuid.UUID,
) (string, error) {
	tkn, err := tokens.NewUploadFileToken(
		tokens.GenerateUploadFilesJwtRequest{
			OwnerAccountID:  OwnerAccountID,
			UploadSessionID: UploadSessionID,
			ResourceID:      OwnerAccountID.String(),
			Resource:        ProfileResource,
			Issuer:          ProfilesActor,
			Audience:        []string{ProfilesActor},
			Ttl:             m.profileMediaUploadTTL,
		}, m.uploadSK,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate upload profile media token, cause: %w", err)
	}

	return tkn, nil
}
