package bucket

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3BucketService struct {
	client     *s3.Client
	bucketName string
}

func NewS3BucketService(client *s3.Client, bucket string) *S3BucketService {
	return &S3BucketService{client: client, bucketName: bucket}
}

func (service S3BucketService) Upload(ctx context.Context, key string, file io.Reader) error {
	manager := manager.NewUploader(service.client, func(u *manager.Uploader) {
		u.PartSize = 10 * 1024 * 1024
	})

	_, err := manager.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(service.bucketName),
		Key:    aws.String(key),
		Body:   file,
	})

	return err
}

func (service S3BucketService) Get(ctx context.Context, key string) ([]byte, error) {
	partSize := int64(10 * 1024 * 1024)
	downloader := manager.NewDownloader(service.client, func(d *manager.Downloader) {
		d.PartSize = partSize
	})
	buffer := manager.NewWriteAtBuffer(make([]byte, partSize))

	_, err := downloader.Download(ctx, buffer, &s3.GetObjectInput{
		Bucket: aws.String(service.bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
