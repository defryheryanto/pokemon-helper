package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/defryheryanto/pokemon-helper/config"
	_ "github.com/golang-migrate/migrate/source"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"
)

func main() {
	config.Load()
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})))

	downFlag := flag.Bool("down", false, "database migration down")
	flag.Parse()

	dsn := config.DatabaseConnectionString()
	if dsn == "" {
		slog.Error("Database connection string is empty")
		return
	}
	slog.Info("Opening database connection to " + dsn)
	db, err := sql.Open("postgres", dsn)
	slog.Info("Database connection opened.")
	if err != nil {
		slog.Error("error opening migration database", "error", err)
		return
	}

	slog.Info("Generating postgres instance...")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		slog.Error("error generating postgres instance", "error", err)
		return
	}
	slog.Info("Postgres instance generated.")

	slog.Info("Opening migration files...")
	fsrc, err := (&file.File{}).Open("file://database/migrations")
	if err != nil {
		slog.Error("error opening migration files", "error", err)
		return
	}
	slog.Info("Migration files opened.")

	slog.Info("Creating migration instance...")
	m, err := migrate.NewWithInstance("file", fsrc, "postgres", driver)
	if err != nil {
		slog.Error("error creating migration instance", "error", err)
		return
	}
	slog.Info("Migration instance created.")

	if *downFlag {
		slog.Info("Rolling back last migration...")
		if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
			slog.Error("error rolling back last migration", "error", err)
			return
		}
		version, _, _ := m.Version()

		slog.Info("Rollback complete.", slog.String("current_version", fmt.Sprintf("%d", version)))
	} else {
		slog.Info("Migrating migration...")
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			slog.Error("error migrating migration", "error", err)
			return
		}
		slog.Info("Migration applied.")
	}
}
