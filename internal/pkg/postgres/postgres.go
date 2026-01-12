package postgres

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func New(databaseConnect string) (*Database, error) {
	db, err := sql.Open("postgres", databaseConnect)
	if err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil
}

func (d *Database) Close() {
	if err := d.db.Close(); err != nil {
		slog.Error("close postgres connection failed", "err", err)
	}
}
