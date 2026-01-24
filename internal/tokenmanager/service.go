package tokenmanager

import (
	"time"

	"github.com/google/uuid"
	"github.com/netbill/restkit/tokens"
)

type Manager struct {
	issuer   string
	uploadSK string

	config Config
}

type Config struct {
	UploadProfileAvatarScope string
	UploadProfileAvatarTtl   time.Duration
}

func New(issuer, uploadSK string, cfg Config) Manager {
	return Manager{
		issuer:   issuer,
		uploadSK: uploadSK,
		config:   cfg,
	}
}

func (m Manager) NewUploadProfileAvatarToken(
	sessionID uuid.UUID,
) (string, error) {
	return tokens.NewUploadFileToken(
		tokens.GenerateUploadFilesJwtRequest{
			SessionID: sessionID,
			Issuer:    m.issuer,
			Audience:  []string{m.issuer},
			Scope:     m.config.UploadProfileAvatarScope,
			Ttl:       m.config.UploadProfileAvatarTtl,
		}, m.uploadSK)
}
