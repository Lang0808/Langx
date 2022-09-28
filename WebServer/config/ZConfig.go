package config

import "github.com/spf13/viper"

func init() {
	viper.SetConfigFile("./config/.env")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Cannot read configurations")
	}

	Instance = &ZConfig{}
}

var Instance Config

type ZConfig struct {
}

func (z *ZConfig) GetConfig(key string) interface{} {
	return viper.Get(key)
}
