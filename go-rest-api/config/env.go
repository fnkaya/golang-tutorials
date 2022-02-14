package config

import "os"

const (
	EnvKey   = "APP_ENV"
	EnvProd  = "prod"
	EnvLocal = "local"
)

func Env() string {
	return GetEnv(EnvKey, EnvLocal)
}

func GetEnv(key, def string) string {
	env, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return env
}
