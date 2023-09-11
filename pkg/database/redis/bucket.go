package redis

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/SergeyCherepiuk/videohosting/domain"
)

type ctxTtlKey string

var CtxTtlKey = ctxTtlKey("ttl")

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
	if err := jsonValue.Err(); err != nil {
		return service.getFromOtherBucketAndCache(ctx, key)
	}

	value := bytesWithContentType{}
	if err := json.Unmarshal([]byte(jsonValue.Val()), &value); err != nil {
		return service.getFromOtherBucketAndCache(ctx, key)
	}

	return value.Bytes, value.ContentType, nil
}

func (service BucketService) Put(ctx context.Context, key, contentType string, file io.Reader) error {
	if ttl, ok := ctx.Value(CtxTtlKey).(time.Duration); ok {
		reader := io.TeeReader(file, &bytes.Buffer{})
		if bytes, err := io.ReadAll(reader); err == nil {
			service.cache(ctx, key, bytes, contentType, ttl)
		}
	}

	return service.otherBucket.Put(ctx, key, contentType, file)
}

func (service BucketService) getFromOtherBucketAndCache(ctx context.Context, key string) ([]byte, string, error) {
	bytes, contentType, err := service.otherBucket.Get(ctx, key)
	if ttl, ok := ctx.Value(CtxTtlKey).(time.Duration); ok {
		service.cache(ctx, key, bytes, contentType, ttl)
	}
	return bytes, contentType, err
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
