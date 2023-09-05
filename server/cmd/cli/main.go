package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SergeyCherepiuk/videohosting/mocks"
	"github.com/SergeyCherepiuk/videohosting/pkg/database/postgres"
	"github.com/SergeyCherepiuk/videohosting/pkg/database/redis"
	"github.com/SergeyCherepiuk/videohosting/pkg/http"
	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
	postgres.MustConnect()
	redis.MustConnect()
}

func main() {
	videoService := postgres.NewVideoService()

	// Amazon S3 Bucket
	// config, err := config.LoadDefaultConfig(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// client := s3.NewFromConfig(config)
	// s3BucketService := aws.NewBucketService(client, os.Getenv("S3_BUCKET_NAME"))

	// Redis Bucket
	// redisBucketService := redis.NewBucketService(s3BucketService)

	// Mocks
	mockBucketService := mocks.NewMockBucketService()

	e := http.Router{
		VideoService: videoService,
		Bucket:       mockBucketService,
	}.Build()
	e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
