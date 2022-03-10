# golang-template-microservice
### (Work in progress)

Sample RESTful microservice to work with products and categories in a small shop.
Has a basic functionality:
- Basic Logging
- Postgres db connection and migrations
- Auto generating Swagger docs
- Business logic
- Configuring using envs

TODO:
- Dockerfile
- Basic Authentication
- Improve logging
- Templates for async jobs
- Cron scheduler
- Profiling with pprof

To get the Swagger page go to: `/api/v1/swagger/index.html`

| Var name                    | Var description                                                                                                              | Default value |
|-----------------------------|------------------------------------------------------------------------------------------------------------------------------|---------------|
| GIN_MODE                    | Run mode for Gin framework. For more info visit the Gin repository.                                                          | debug         |
| LOG_LEVEL                   | Logging level.                                                                                                               | DEBUG         |
| LISTEN_ADDRESS              | Services' port.                                                                                                              | 8080          |
| POSTGRES_HOST               |                                                                                                                              | localhost     |
| POSTGRES_PORT               |                                                                                                                              | 5432          |
| POSTGRES_DB                 | Postgres database. Should be created in advance. After the service started, migrations will be applied (two tables created). | store         |
| POSTGRES_USERNAME           |                                                                                                                              | postgres      |
| POSTGRES_PASSWORD           |                                                                                                                              | postgres      |
| POSTGRES_SSL_MODE           |                                                                                                                              | disable       |
| POSTGRES_CONNECTION_TIMEOUT |                                                                                                                              | 10            |