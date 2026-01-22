package main

import (
	"database/sql"

	"github.com/defryheryanto/pokemon-helper/config"
	"github.com/defryheryanto/pokemon-helper/internal/logger"
	_ "github.com/lib/pq"
)

func setupDB() *sql.DB {
	dsn := config.DatabaseConnectionString()
	if dsn == "" {
		panic("ENV DATABASE_CONNECTION_STRING is empty")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error("error opening database connection", err)
		panic(err)
	}

	return db
}
