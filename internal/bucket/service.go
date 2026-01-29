package bucket

import (
	"context"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"time"
)

type Bucket struct {
	s3 awsxs3
}

func New(awsx3 awsxs3) Bucket {
	return Bucket{
		s3: awsx3,
	}
}

type awsxs3 interface {
	PresignPut(
		ctx context.Context,
		key string,
		ttl time.Duration,
	) (uploadURL, getUrl string, error error)

	GetObjectRange(
		ctx context.Context,
		key string,
		bytes int64,
	) (body io.ReadCloser, size int64, err error)
	CopyObject(ctx context.Context, tmplKey, finalKey string) (string, error)
	DeleteObject(ctx context.Context, key string) error
}
