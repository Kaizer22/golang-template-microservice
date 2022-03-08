package db

import "database/sql"

const (
	PgMigrationsPath = "file://db/migration/postgres"
)

type ConnectionProvider interface {
	Connection() *sql.DB
	IsConnected() (bool, error)
	Migrate(migrationPath string) error
}
