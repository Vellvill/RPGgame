package config

import (
	"github.com/joho/godotenv"
	"os"
	"sync"
)

type Config struct {
	Logger struct {
		LogLevel string
	}
	Db struct {
		DbHost string
		DbPort string
		DbUser string
		DbPass string
		DbName string
	}
	Service struct {
		Host string
		Port string
	}
}

var (
	once     sync.Once
	instance *Config
	err      error
)

func NewConfig() (*Config, error) {

	once.Do(func() {
		err = godotenv.Load("./internal/config/.env")
		if err != nil {
			return
		}
		instance = &Config{
			Logger: struct {
				LogLevel string
			}{
				LogLevel: parseENV("LOG_LEVEL", "debug"),
			},
			Db: struct {
				DbHost string
				DbPort string
				DbUser string
				DbPass string
				DbName string
			}{
				DbHost: parseENV("DB_HOST", "db"),
				DbPort: parseENV("DB_PORT", "6543"),
				DbUser: parseENV("DB_USER", ""),
				DbPass: parseENV("DB_PASS", ""),
				DbName: parseENV("DB_NAME", ""),
			},
			Service: struct {
				Host string
				Port string
			}{
				Host: parseENV("SERVICE_HOST", "service"),
				Port: parseENV("SERVICE_PORT", "40000")},
		}
	})
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func parseENV(env, defaultValue string) string {

	if value, exists := os.LookupEnv(env); exists {
		return value
	}
	return defaultValue
}
