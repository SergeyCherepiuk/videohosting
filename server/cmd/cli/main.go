package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SergeyCherepiuk/videohosting/pkg/bucket"
	"github.com/SergeyCherepiuk/videohosting/pkg/http"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Amazon S3
	// config, err := config.LoadDefaultConfig(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// client := s3.NewFromConfig(config)
	// s3BucketService := bucket.NewS3BucketService(client, os.Getenv("S3_BUCKET_NAME"))

	// Mocks
	mockBucketService := bucket.NewMockBucketService()

	e := http.Router{
		Bucket: mockBucketService,
	}.Build()
	e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
