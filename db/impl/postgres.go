package impl

const (
	PostgresHostEnv     = "POSTGRES_HOST"
	PostgresPortEnv     = "POSTGRES_PORT"
	PostgresDatabaseEnv = "POSTGRES_DB"
	PostgresUsernameEnv = "POSTGRES_USERNAME"
	PostgresPasswordEnv = "POSTGRES_PASSWORD"
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
}

type pgConnectionProvider struct {
	options pgConnectionOptions
}

func NewPgConnectionProvider() {

}

//func GetPgConnection() *pgConnectionProvider {
//	connInit.Do(func() {
//
//		if conn, err := dbConnection.IsConnected(); !conn {
//			logging.Error("Cannot establish connection to a DB")
//			panic(err)
//		}
//	})
//	return dbConnection
//}
//
//func createSchema(db *pg.DB) error {
//
//	return nil
//}
//
//const ConnectionTestQuery = "SELECT 1"
//func (p *PgConnection) IsConnected() (bool, error) {
//	if _, err := p.connection.Exec(ConnectionTestQuery); err != nil {
//		return false, err
//	}
//	return true, nil
//}
//
//func (p *PgConnection) Migrate() error {
//	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
//		p.options.User,
//		p.options.Password,
//		p.options.Addr,
//		p.options.Database)
//	fmt.Println(connStr)
//	logging.InfoFormat("Start migration with connection string %s", connStr)
//
//	m, err := migrate.New(MigrationsPath, connStr)
//	if err != nil {
//		return err
//	}
//
//	err = m.Up()
//	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
//		return err
//	}
//
//	return nil
//}
//
//func (p *PgConnection) Connection() *pg.DB {
//	return p.connection.Conn().
//}
//
//func connectionOptions() *pg.Options  {
//	host := utils.GetEnv(repository.PostgresHostEnv, "localhost")
//	port := utils.GetEnv(repository.PostgresPortEnv, "5432")
//	//logging.InfoFormat()
//	opts := &pg.Options{
//		Addr:                  fmt.Sprintf("%s:%s", host, port),
//		User:                  utils.GetEnv(repository.PostgresUsernameEnv, "postgres"),
//		Password:              utils.GetEnv(repository.PostgresPasswordEnv, "postgres"),
//		Database:              utils.GetEnv(repository.PostgresDatabaseEnv, "geo_object_provider"),
//		//ApplicationName:       "",
//		//TLSConfig:             nil,
//		//DialTimeout:           0,
//		//ReadTimeout:           0,
//		//WriteTimeout:          0,
//		//MaxRetries:            0,
//		//RetryStatementTimeout: false,
//		//MinRetryBackoff:       0,
//		//MaxRetryBackoff:       0,
//		//PoolSize:              0,
//		//MinIdleConns:          0,
//		//MaxConnAge:            0,
//		//PoolTimeout:           0,
//		//IdleTimeout:           0,
//		//IdleCheckFrequency:    0,
//	}
//	logging.InfoFormat("PostgreSQL connection params: %v", &opts)
//	return opts
//}
