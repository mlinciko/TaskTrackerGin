package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	JWTSecretKey string
}

var AppConfig *Config

func LoadConfig() {
	// Загружаем .env файл, если есть
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Println("No .env file found, reading config from environment variables")
	}

	AppConfig = &Config{
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", "mlinciko"),
		DBName:       getEnv("DB_NAME", "postgres"),
		JWTSecretKey: getEnv("JWT_SECRET", ""),
	}
}

// getEnv читает переменную окружения или возвращает значение по умолчанию
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// Формирует DSN для подключения к PostgreSQL
func (c *Config) GetPostgresDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName,
	)
}

func (c *Config) GetTestPostgresDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName+"_test",
	)
}
