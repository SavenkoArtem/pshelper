package logger

import (
	"os"

	"github.com/SavenkoArtem/pshelper/configs"
	"github.com/rs/zerolog"
)

func InitLogger(config *configs.LogConfig) *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.Level(config.Lever))
	var logger zerolog.Logger
	if config.Format == "json" {
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "15:04:05", // Короткий формат времени

			FormatLevel: func(i interface{}) string {
				var l string
				if ll, ok := i.(string); ok {
					switch ll {
					case "debug":
						l = "\x1b[34mDBG\x1b[0m" // Синий
					case "info":
						l = "\x1b[32mINF\x1b[0m" // Зеленый
					case "warn":
						l = "\x1b[33mWRN\x1b[0m" // Желтый
					case "error":
						l = "\x1b[31mERR\x1b[0m" // Красный
					default:
						l = ll
					}
				}
				return l
			},

			FormatMessage: func(i interface{}) string {
				return "\x1b[1m" + i.(string) + "\x1b[0m" // Жирный текст
			},
		}
		logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	}
	return &logger
}
