package connect

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewConnection() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:example@0.0.0.0:5432/postgres")
	if err != nil {
		return nil, err
	}
	return conn, nil
}
