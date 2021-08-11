package main

import (
	"log"

	"github.com/StartloJ/lab-redis-k8s/internal/config"
	"github.com/StartloJ/lab-redis-k8s/pkg/client"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("REDEV_ENDPOINT", "localhost:6379")
	viper.SetDefault("REDEV_CLUSTER", false)
	viper.SetDefault("REDEV_PASSWORD", "")
	viper.SetDefault("REDEV_INIT_DB", 0)
	viper.SetDefault("REDEV_ROUND", 10)
}

func main() {
	log.Println("Start demo app")
	config.Init()

	if viper.GetBool("cluster") {
		log.Printf("Now, we does not support cluster mode with REDEV_CLUSTER: %s", viper.GetString("cluster"))
	} else if !viper.GetBool("cluster") {
		client.Execution()
	}
}
