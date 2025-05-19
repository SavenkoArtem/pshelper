package configs

import (
	"log"

	"github.com/joho/godotenv"
)

// AppConfig содержит все конфигурации приложения
type AppConfig struct {
	Server ServerConfig
	Db     DbConfig
	Logger LogConfig
}

// ServerConfig содержит конфигурации сервера
type ServerConfig struct {
	Port string // Порт, на котором будет работать сервер
}

// DbConfig содержит конфигурации базы данных
type DbConfig struct {
	Dsn string
}

type LogConfig struct {
	Lever  int
	Format string
}

// LoadConfig загружает конфигурацию из .env файла и переменных окружения
func LoadConfig() *AppConfig {
	// Загружаем .env файл
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Парсим порт из переменных окружения
	return &AppConfig{
		Server: ServerConfig{
			Port: getString("SERVER_PORT", "8080"),
		},
		Db: DbConfig{
			Dsn: getString("DSN", ""),
		},
		Logger: LogConfig{
			Lever:  getInt("LOG_LEVEL", 1),
			Format: getString("LOG_FORMAT", "json"),
		},
	}
}
