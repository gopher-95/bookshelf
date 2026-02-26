package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	//конфиг базы данных
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	//конфиг сервера
	ServerPort string
}

// функция загрузки конфиг файла
func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	config := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "bookshelf_db"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}

	if config.DBPassword == "" {
		return nil, fmt.Errorf("DB_PASSWORD is required")
	}

	log.Println("Config loaded successfully")

	return config, nil

}

// функция чтения конфиг файла
func getEnv(key, defaulValue string) string {
	//читает .env файл и возвращает значение ключа
	value := os.Getenv(key)
	if value == "" {
		return defaulValue
	}

	return value
}

// функция возвращает строку подключения к бд
func (c *Config) DBConnectionString() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode)
}

// функция возвращает номер порта
func (c *Config) ServerPortString() string {
	return fmt.Sprintf(":%s", c.ServerPort)
}
