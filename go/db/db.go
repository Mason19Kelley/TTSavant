package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)


func GetDBConnection() (*pgx.Conn, error) {
    conn, err := pgx.Connect(context.Background(), "postgres://postgres:password@localhost:5432/postgres")
    if err != nil {
        return nil, err
    }
    return conn, nil
}

