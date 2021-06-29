package main

import (
	"context"
	"log"

	"github.com/StartloJ/lab-redis-k8s/internal/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var ctx = context.Background()

func init() {
	viper.SetDefault("REDEV_ENDPOINT", "localhost:6379")
	viper.SetDefault("REDEV_PASSWORD", "")
	viper.SetDefault("REDEV_INIT_DB", 0)
}

func main() {
	log.Println("Start demo app")

	config.Init()
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("endpoint"),
		Password: viper.GetString("password"), // no password set
		DB:       viper.GetInt("init_db"),     // use default DB
	})
	pong, err := rdb.Ping(ctx).Result()
	log.Println(pong, err)
}
