package config

type Config interface {
	GetConfig(key string) interface{}
}
