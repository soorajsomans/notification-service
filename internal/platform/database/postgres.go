package database

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func NewPostgres() (*sqlx.DB, error) {
	dsn := "host=localhost " +
		"port=5432 " +
		"user=postgres " +
		"password=postgres " +
		"dbname=notification_db " +
		"sslmode=disable"

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
