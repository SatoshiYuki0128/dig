package config

import (
	"github.com/spf13/viper"
)

var Config *viper.Viper

func InitConfig() error {
	Config = viper.New()
	Config.SetConfigName("config")
	Config.SetConfigType("json")
	Config.AddConfigPath("./config")

	err := Config.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
