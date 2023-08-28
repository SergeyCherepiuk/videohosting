package bucket

import (
	"context"
	"io"
)

type BucketService interface {
	Upload(ctx context.Context, key string, file io.Reader) error
	Get(ctx context.Context, key string) ([]byte, error)
}
