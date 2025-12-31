package config

import (
	"errors"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	configInstance *Config
	once           sync.Once
)

type Postgres struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	SSL_MODE    string
	TIMEZONE    string
}

type Config struct {
	APP_ENV string
	PORT    string

	BCRYPT_COST string
	JWT_SECRET  string

	POSTGRES Postgres
}

// LoadConfig 함수는 Config 인스턴스를 반환
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("Cannot Find .env File")
	}

	return &Config{
		APP_ENV:     getEnv("APP_ENV", "development"),
		PORT:        getEnv("PORT", "6789"),
		BCRYPT_COST: getEnv("BCRYPT_COST", "5"),
		JWT_SECRET:  getEnv("JWT_SECRET", "secret"),
		POSTGRES: Postgres{
			DB_HOST:     getEnv("DB_HOST", "localhost"),
			DB_USER:     getEnv("DB_USER", "postgres"),
			DB_PASSWORD: getEnv("DB_PASSWORD", "postgres"),
			DB_NAME:     getEnv("DB_NAME", "action_manager"),
			DB_PORT:     getEnv("DB_PORT", "5432"),
			SSL_MODE:    getEnv("SSL_MODE", "disable"),
			TIMEZONE:    getEnv("TIMEZONE", "Asia/Shanghai"),
		},
	}, nil
}

// GetConfig 함수는 Config 인스턴스를 반환
func GetConfig() *Config {
	once.Do(func() {
		configInstance, _ = LoadConfig()
	})
	return configInstance
}

// getEnv 함수는 환경 변수에서 값을 가져오고, 없으면 기본값을 반환
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
