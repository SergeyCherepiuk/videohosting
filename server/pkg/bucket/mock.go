package bucket

import (
	"context"
	"io"
	"time"
)

type MockBucketService struct{}

func NewMockBucketService() *MockBucketService {
	return &MockBucketService{}
}

func (service MockBucketService) Get(ctx context.Context, key string) ([]byte, string, error) {
	return make([]byte, 32*1024*1024), "text/plain", nil
}

func (service MockBucketService) Upload(ctx context.Context, key, contentType string, file io.Reader) error {
	buffer := make([]byte, 10*1024*1024)
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
