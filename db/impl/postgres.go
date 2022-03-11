package impl

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"main/db"
	"main/logging"
	"main/utils"
)

const (
	PostgresHostEnv              = "POSTGRES_HOST"
	PostgresPortEnv              = "POSTGRES_PORT"
	PostgresDatabaseEnv          = "POSTGRES_DB"
	PostgresUsernameEnv          = "POSTGRES_USERNAME"
	PostgresPasswordEnv          = "POSTGRES_PASSWORD"
	PostgresSSLModeEnv           = "POSTGRES_SSL_MODE"
	PostgresConnectionTimeoutEnv = "POSTGRES_CONNECTION_TIMEOUT"
)

type pgConnectionOptions struct {
	PgHost     string
	PgPort     string
	PgUsername string
	PgPassword string
	PgDatabase string
	// PgConnectionTimeout timeout in seconds
	PgConnectionTimeout int64
	PgSslMode           string

	PgConnectionString string
}

type pgConnectionProvider struct {
	options pgConnectionOptions
	pgDb    *sql.DB
}

func NewPgConnectionProvider() (db.ConnectionProvider, error) {
	options := pgConnectionOptions{
		PgHost:              utils.GetEnv(PostgresHostEnv, "localhost"),
		PgPort:              utils.GetEnv(PostgresPortEnv, "5432"),
		PgUsername:          utils.GetEnv(PostgresUsernameEnv, "postgres"),
		PgPassword:          utils.GetEnv(PostgresPasswordEnv, "postgres"),
		PgDatabase:          utils.GetEnv(PostgresDatabaseEnv, "promo"),
		PgConnectionTimeout: utils.GetEnvInt(PostgresConnectionTimeoutEnv, 10),
		PgSslMode:           utils.GetEnv(PostgresSSLModeEnv, "disable"),
	}
	options.PgConnectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		options.PgUsername, options.PgPassword,
		options.PgHost, options.PgPort,
		options.PgDatabase, options.PgSslMode)

	dbConnection, err := sql.Open("postgres", options.PgConnectionString)
	if err != nil {
		logging.FatalFormat("cannot open db connection %s",
			options.PgConnectionString, err)
		return nil, err
	}
	logging.DebugFormat("opened a new pg connection - %s", options.PgConnectionString)
	provider := pgConnectionProvider{
		options: options,
		pgDb:    dbConnection,
	}
	if c, err := provider.IsConnected(); err != nil || !c {
		logging.FatalFormat("cannot establish a connection to %s", options.PgConnectionString)
		return nil, err
	}

	return provider, nil
}

const ConnectionTestQuery = "SELECT 1"

func (p pgConnectionProvider) IsConnected() (bool, error) {
	if _, err := p.pgDb.Exec(ConnectionTestQuery); err != nil {
		return false, err
	}
	return true, nil
}

func (p pgConnectionProvider) Migrate(migrationPath string) error {
	//connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
	//	p.options.PgUsername,
	//	p.options.PgPassword,
	//	p.options.PgHost,
	//	p.options.PgDatabase)
	//fmt.Println(connStr)
	//logging.InfoFormat("Start migration with connection string %s", connStr)

	//p.createDbIfNeeded()
	db, err := sql.Open("postgres", p.options.PgConnectionString)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		"postgres", driver)

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	db.Close()
	return nil
}

func (p pgConnectionProvider) Connection() *sql.DB {
	return p.pgDb
}

func (p pgConnectionProvider) createDbIfNeeded() error {
	if _, err := p.pgDb.Exec(
		fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", p.options.PgDatabase),
	); err != nil {
		return err
	}
	return nil
}
