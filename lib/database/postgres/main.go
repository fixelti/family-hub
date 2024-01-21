package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	*pgx.Conn
}

func New(ctx context.Context ,dsn string) Database {
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}
	
	return Database{conn}
}

func ScanInStruct[T any](rows pgx.Rows) (*T, error) {
	return pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[T])
}