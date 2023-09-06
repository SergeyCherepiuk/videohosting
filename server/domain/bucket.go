package domain

import (
	"context"
	"io"
)

type BucketService interface {
	Get(ctx context.Context, key string) ([]byte, string, error)
	Put(ctx context.Context, key, contentType string, file io.Reader) error
}
