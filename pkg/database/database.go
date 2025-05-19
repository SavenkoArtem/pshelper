package database

import (
	"context"
	"log"

	"github.com/SavenkoArtem/pshelper/configs"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

func CreateDbPool(config *configs.AppConfig, logger *zerolog.Logger) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.Db.Dsn)
	if err != nil {
		logger.Error().Msg(ErrNotConnected)
		log.Println(ErrNotConnected)
		panic(ErrNotConnectedDB)
	}

	return dbpool
}
