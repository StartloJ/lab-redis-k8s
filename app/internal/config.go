package internal

import "github.com/spf13/viper"

// Initial function for prepared ENV config.
func init() {
	viper.SetEnvPrefix("redev")
	viper.AutomaticEnv()

	viper.SetDefault("REDEV_ENDPOINT", "localhost:6379")
	viper.SetDefault("REDEV_PASSWORD", "")
}
