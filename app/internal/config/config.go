package config

import "github.com/spf13/viper"

// Initial function for prepared ENV config.
func Init() {
	viper.SetEnvPrefix("redev")
	viper.AutomaticEnv()

}
