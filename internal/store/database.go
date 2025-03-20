package store

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
	"io/fs"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("pgx",
		"host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("db open error: %w", err)
	}
	fmt.Println("Connected to database")
	return db, nil
}

func MigrateFS(db *sql.DB, migrationsFS fs.FS, dir string) error {
	goose.SetBaseFS(migrationsFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()

	return Migrate(db, dir)
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("Migrate:  %w", err)
	}

	err = goose.Up(db, dir)
	fmt.Println("Migration complete:", dir)
	if err != nil {
		return fmt.Errorf("Goose up:  %w", err)
	}
	fmt.Println("Migration complete")
	return nil
}
