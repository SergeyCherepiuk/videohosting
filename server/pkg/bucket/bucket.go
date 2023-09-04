package bucket

import (
	"context"
	"io"
)

type BucketService interface {
	Get(ctx context.Context, key string) ([]byte, string, error)
	Upload(ctx context.Context, key, contentType string, file io.Reader) error
}
