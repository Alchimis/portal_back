package cmd

import (
	stdsql "database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"path/filepath"
)

func Migrate(config Config) error {
	conn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", config.DBUser, config.DBPassword, config.DBHost, config.DBName)
	db, err := stdsql.Open("postgres", conn)
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	migrations, err := filepath.Abs("./documentation/db/migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:%s", migrations),
		config.DBName, driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}
	return err
}
