package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func New(databaseConnect string) (*Database, error) {
	db, err := sql.Open("postgres", databaseConnect)
	if err != nil {
		return nil, wrapErr(err)
	}

	if err := db.Ping(); err != nil {
		return nil, wrapErr(err)
	}

	return &Database{
		db: db,
	}, nil
}

func (d *Database) Close() error {
	if err := d.db.Close(); err != nil {
		return wrapErr(err)
	}

	return nil
}

func wrapErr(err error) error {
	return fmt.Errorf("postgres: %v", err)
}
