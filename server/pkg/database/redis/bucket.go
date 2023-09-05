package redis

import (
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/SergeyCherepiuk/videohosting/domain"
)

type BucketService struct {
	otherBucket domain.BucketService
}

func NewBucketService(otherBucket domain.BucketService) *BucketService {
	return &BucketService{otherBucket: otherBucket}
}

type bytesWithContentType struct {
	Bytes       []byte `json:"bytes"`
	ContentType string `json:"content_type"`
}

func (service BucketService) Get(ctx context.Context, key string) ([]byte, string, error) {
	jsonValue := db.Get(ctx, key)
	if jsonValue.Err() != nil {
		bytes, contentType, err := service.otherBucket.Get(ctx, key)
		service.cache(ctx, key, bytes, contentType, 7*24*time.Hour)
		return bytes, contentType, err
	}

	value := bytesWithContentType{}
	if err := json.Unmarshal([]byte(jsonValue.Val()), &value); err != nil {
		bytes, contentType, err := service.otherBucket.Get(ctx, key)
		service.cache(ctx, key, bytes, contentType, 7*24*time.Hour)
		return bytes, contentType, err
	}

	return value.Bytes, value.ContentType, nil
}

func (service BucketService) Upload(ctx context.Context, key, contentType string, file io.Reader) error {
	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	service.cache(ctx, key, bytes, contentType, 7*24*time.Hour)

	return service.otherBucket.Upload(ctx, key, contentType, file)
}

func (service BucketService) cache(
	ctx context.Context, key string, bytes []byte, contentType string, ttl time.Duration,
) {
	if value, err := json.Marshal(bytesWithContentType{
		Bytes:       bytes,
		ContentType: contentType,
	}); err == nil {
		db.Set(ctx, key, value, ttl)
	}
}
