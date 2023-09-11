package mocks

import (
	"context"
	"io"
	"time"
)

type MockBucketService struct{}

func NewMockBucketService() *MockBucketService {
	return &MockBucketService{}
}

var partSize = int64(10 * 1024 * 1024)

func (service MockBucketService) Get(ctx context.Context, key string) ([]byte, string, error) {
	return make([]byte, 5*partSize), "text/plain", nil
}

func (service MockBucketService) Put(ctx context.Context, key, contentType string, file io.Reader) error {
	buffer := make([]byte, partSize)
	for {
		time.Sleep(100 * time.Millisecond)
		_, err := file.Read(buffer)
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
	}
}
