package config

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/dotnev"
)

func LoadConfig() {
	config.WithOptions(config.ParseEnv)

	err := dotnev.Load("./", ".env")
	if err != nil {
		panic("Cannot load .env")
	}

	config.LoadOSEnv([]string{
		"APP_NAME",
		"DB_HOST",
		"DB_NAME",
		"DB_USER",
		"DB_PASS",
		"DB_PORT",
		"TIME_ZONE",
	}, true)

}

func Get(key string) interface{} {
	return dotnev.Get(key)
}
