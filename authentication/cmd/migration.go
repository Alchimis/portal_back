package cmd

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"path/filepath"
)

func Migrate(config Config) error {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, 5432, config.DBUser, config.DBPassword, config.DBName)
	//db, err := stdsql.Open("postgres", conn)

	//if err != nil {
	//	fmt.Println("1")
	//	return err
	//}

	//driver, err := postgres.WithInstance(db, &postgres.Config{})
	//if err != nil {
	//	fmt.Println("2")
	//	return err
	//}

	migrations, err := filepath.Abs("authentication/db/migrations")
	m, err := migrate.New(
		"file://app/authentication/db/migrations",
		conn)
	fmt.Println(migrations)
	if err != nil {
		fmt.Println("3")
		return err
	}

	//m, err := migrate.NewWithDatabaseInstance(
	//	fmt.Sprintf("file:///%s", migrations),
	//	config.DBName, driver)
	if err != nil {
		fmt.Println("4")
		return err
	}

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("5")
		return nil
	}
	fmt.Println("6")

	return err
}
