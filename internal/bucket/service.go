package bucket

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type Bucket struct {
	awsx3 awsx3

	config Config
}

type Config struct {
	ProfileAvatarUploadTTL  time.Duration
	ProfileAvatarMaxLength  int64
	ProfileAvatarAllowedExt []string
}

func New(awsx3 awsx3, cfg Config) Bucket {
	return Bucket{
		awsx3:  awsx3,
		config: cfg,
	}
}

type awsx3 interface {
	PresignPut(
		ctx context.Context,
		key string,
		contentLength int64,
		ttl time.Duration,
	) (uploadURL, getUrl string, error error)

	HeadObject(ctx context.Context, key string) (*s3.HeadObjectOutput, error)
	CopyObject(ctx context.Context, tmplKey, finalKey string) (string, error)
	DeleteObject(ctx context.Context, key string) error
}

func CreateTempProfileAvatarKey(accountID, sessionID uuid.UUID) string {
	return fmt.Sprintf("profile/avatar/%s/temp/%s", accountID, sessionID)
}

func CreateProfileAvatarKey(accountID uuid.UUID) string {
	return fmt.Sprintf("profile/avatar/%s", accountID)
}

func (r Bucket) GetPreloadLinkForUpdateProfileAvatar(
	ctx context.Context,
	accountID, sessionID uuid.UUID,
) (uploadURL, getUrl string, error error) {
	uploadURL, getURL, err := r.awsx3.PresignPut(
		ctx,
		CreateTempProfileAvatarKey(accountID, sessionID),
		r.config.ProfileAvatarMaxLength,
		r.config.ProfileAvatarUploadTTL,
	)
	if err != nil {
		return "", "", err
	}

	return uploadURL, getURL, nil
}

func (r Bucket) AcceptUpdateProfileAvatar(
	ctx context.Context,
	accountID, sessionID uuid.UUID,
) (string, error) {
	return r.awsx3.CopyObject(ctx,
		CreateTempProfileAvatarKey(accountID, sessionID),
		CreateProfileAvatarKey(accountID),
	)
}

func (r Bucket) CheckProfileAvatarExtension(link string) (bool, error) {
	head, err := r.awsx3.HeadObject(context.TODO(), link) // или просто HeadObject внутри bucket
	if err != nil {
		return false, err
	}

	ct := ""
	if head.ContentType != nil {
		ct = *head.ContentType
	}

	switch ct {
	case "image/png", "image/jpeg":
		// ok
	default:
		return false, fmt.Errorf("content type is invalid, allowed only: image/png, image/jpeg")
	}

	return true, nil
}

func (r Bucket) CancelUpdateProfileAvatar(
	ctx context.Context,
	accountID, sessionID uuid.UUID,
) error {
	err := r.awsx3.DeleteObject(ctx, CreateTempProfileAvatarKey(accountID, sessionID))
	if err != nil {
		return err
	}

	return nil
}
