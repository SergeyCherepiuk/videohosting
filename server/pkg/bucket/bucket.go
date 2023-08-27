package bucket

import (
	"context"
)

type BucketService interface {
	Upload(ctx context.Context, key string, file []byte) error
	Get(ctx context.Context, ket string) ([]byte, error)
}
