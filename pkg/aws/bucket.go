package aws

import (
	"context"
	"io"

	"github.com/SergeyCherepiuk/videohosting/pkg/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type BucketService struct {
	client     *s3.Client
	bucketName string
}

func NewBucketService(client *s3.Client, bucket string) *BucketService {
	return &BucketService{client: client, bucketName: bucket}
}

func (service BucketService) Get(ctx context.Context, key string) ([]byte, string, error) {
	input := s3.GetObjectInput{
		Bucket: aws.String(service.bucketName),
		Key:    aws.String(key),
	}

	output, err := service.client.GetObject(ctx, &input)
	if err != nil {
		return nil, "", err
	}

	downloader := manager.NewDownloader(service.client, func(d *manager.Downloader) {
		d.PartSize = config.PartSize
	})
	buffer := manager.NewWriteAtBuffer([]byte{})
	if _, err := downloader.Download(ctx, buffer, &input); err != nil {
		return nil, "", err
	}

	return buffer.Bytes(), *output.ContentType, nil
}

func (service BucketService) Put(ctx context.Context, key, contentType string, file io.Reader) error {
	manager := manager.NewUploader(service.client, func(u *manager.Uploader) {
		u.PartSize = config.PartSize
	})

	_, err := manager.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(service.bucketName),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
		Body:        file,
	})

	return err
}
