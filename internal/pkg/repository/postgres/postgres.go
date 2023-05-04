package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gookit/config/v2"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
	"os"
)

func NewDB() *bun.DB {
	pDB, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Data()["postgres.host"], config.Data()["postgres.port"], config.Data()["postgres.user"], config.Data()["postgres.password"], config.Data()["postgres.dbname"]))
	if err != nil {
		log.Fatal(err)
	}
	defer pDB.Close()

	db := bun.NewDB(pDB, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	migrateUp(db.DB)

	return db
}

func migrateUp(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", config.Data()[`migration.path`].(string)),
		config.Data()["postgres.dbname"].(string), driver)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		fmt.Println(err)
		os.Exit(1)
	}
}
