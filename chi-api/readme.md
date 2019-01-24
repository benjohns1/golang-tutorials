Start docker postgres container: `docker-compose up`

First time setup test DB: `go run persistence/postgresdb/setup/setup.go`

Run database script: `go build && chi-api`