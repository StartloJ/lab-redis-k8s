package config

import "github.com/spf13/viper"

// Initial function for prepared ENV config.
func init() {
	viper.SetEnvPrefix("redev")
	viper.AutomaticEnv()

}
