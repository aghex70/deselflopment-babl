package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/aghex70/deselflopment-babl/internal/stores/calendar"
	"github.com/aghex70/deselflopment-babl/internal/stores/entry"
	"github.com/aghex70/deselflopment-babl/internal/stores/user"
	"github.com/pressly/goose"
	"gorm.io/gorm"
)

var (
	ErrMigrationFile = errors.New("please provide a valid name for the migration file")
)

const (
	// Migration directory
	migrationDirectory string = "persistence/database/migrations"

	// Migrations table
	migrationTable string = "deselflopment-babl_db_version"
)

func init() {
	goose.SetTableName(migrationTable)
}

func Migrate(db *sql.DB) error {
	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	if err := goose.Run("up", db, migrationDirectory, "sql"); err != nil {
		fmt.Printf("%+v", err)
		return err
	}

	return nil
}

func MakeMigrations(db *sql.DB, filename string) error {
	if filename == "" {
		return ErrMigrationFile
	}

	if err := goose.Run("create", db, migrationDirectory, filename, "sql"); err != nil {
		return err
	}

	return nil
}

func AutoMigrateWithTimestamp(db *gorm.DB) error {
	migrator := db.Migrator()

	if err := migrator.AutoMigrate(&calendar.Calendar{}, &entry.Entry{}, &user.User{}); err != nil {
		return err
	}

	return nil
}
