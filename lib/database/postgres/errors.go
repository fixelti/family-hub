package postgres

import "github.com/jackc/pgx/v5"

var (
	ErrNotFound = pgx.ErrNoRows
)