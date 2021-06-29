package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	log.Println("Start demo app")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "P@ssw0rd", // no password set
		DB:       0,          // use default DB
	})
	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
}
