package config

import (
	"os"
)

type Config struct {
	App AppConfig
	Log LogConfig
}

type AppConfig struct {
	Port string
	Env  string
}

type LogConfig struct {
	Level  string
	Pretty bool
}

func Load() *Config {
	return &Config{
		App: AppConfig{
			Port: getEnv("APP_PORT", "8080"),
			Env:  getEnv("APP_ENV", "development"),
		},
		Log: LogConfig{
			Level:  getEnv("LOG_LEVEL", "debug"),
			Pretty: getEnv("APP_ENV", "development") == "development",
		},
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
