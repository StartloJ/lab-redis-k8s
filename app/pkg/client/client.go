package client

import (
	"context"
	"log"
	"math/rand"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var ctx = context.Background()

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const randInt = 12

func setter(rdb *redis.Client, key string, value interface{}) error {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Panic(err)
		return err
	}
	return nil
}

func getter(rdb *redis.Client, key string) string {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Panic(err)
	}
	return val
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func Execution() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("endpoint"),
		Password: viper.GetString("password"), // no password set
		DB:       viper.GetInt("init_db"),     // use default DB
	})

	for run := 1; run < viper.GetInt("round"); run++ {
		val := RandStringBytes(randInt)
		err := setter(rdb, "key"+strconv.Itoa(run), val)
		if err != nil {
			log.Panicf("Can not add key in DB with %v", err)
		}
		getVal := getter(rdb, "key"+strconv.Itoa(run))
		log.Printf("Add data in redis with value: %s", getVal)
	}
}
