package db

import (
	"btaw/cmd/api/cfg"
	"btaw/logger"
	"context"

	"github.com/jackc/pgx/v5"
)

func Query() error {
	conn, err := pgx.Connect(context.Background(), cfg.DATABASE_URL)
	if err != nil {
		logger.Log.Printf("Unable to connect to database: %v\n", err)
		return err
	}
	defer conn.Close(context.Background())

	var ackConnect string
	err = conn.QueryRow(context.Background(), "select 'DB connection ok'").Scan(&ackConnect)
	if err != nil {
		logger.Log.Printf("QueryRow failed: %v\n", err)
		return err
	}

	logger.Log.Println(ackConnect)

	return nil

}
