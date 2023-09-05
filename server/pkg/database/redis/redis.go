package redis

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var db *redis.Client

func MustConnect() {
	db = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
	})
	if err := db.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}
	db.SAdd(context.Background(), fmt.Sprintf("%s_keys", os.Getenv("REDIS_PREVIEW_PREFIX")))
}
