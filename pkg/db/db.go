package db

import (
	"btaw/cmd/gw/bx/app"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func Query() error {
	conn, err := pgx.Connect(context.Background(), app.DatabaseURL)
	if err != nil {
		log.Error().Err(err).Msg("Unable to connect to database")
		return err
	}
	defer conn.Close(context.Background())

	var ackConnect string
	err = conn.QueryRow(context.Background(), "select 'DB connection ok'").Scan(&ackConnect)
	if err != nil {
		log.Error().Err(err).Msg("QueryRow failed")
		return err
	}

	log.Printf(ackConnect)

	return nil

}
