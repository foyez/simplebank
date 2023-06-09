DB_USER=root
DB_PASSWORD=testpass
DB_NAME=simplebank
DB_URL=postgres://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable
DB_CONTAINER=postgres15

## db_docs: generate dbdocs view
db_docs:
	@echo "generating dbdocs view..."
	dbdocs build doc/db.dbml

## db_schema: convert a dbml file to sql
db_schema:
	@echo "converting a dbml file to sql..."
	dbml2sql doc/db.dbml --postgres -o doc/schema.sql

## postgres: run postgres container
postgres:
	@echo "running postgres container..."
	docker run --name ${DB_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -d postgres:15.2-alpine

## createdb: create a database
createdb:
	@echo "creating ${DB_NAME} database..."
	docker exec -it ${DB_CONTAINER} createdb --username=root --owner=root ${DB_NAME}

## dropdb: delete a database
dropdb:
	@echo "deleting ${DB_NAME} database..."
	docker exec -it ${DB_CONTAINER} dropdb ${DB_NAME}

## create_migration: create migration up & down files
create_migration:
	@echo "creating migration files..."
	migrate create -ext sql -dir db/migration -seq ${name}

## migrateup: apply all up migrations
migrateup:
	@echo "applying all up migrations..."
	migrate -path db/migration -database ${DB_URL} -verbose up

## migratedown: apply all down migrations
migratedown:
	@echo "applying all down migrations..."
	migrate -path db/migration -database ${DB_URL} -verbose down

## sqlc: generate Go code from SQL
sqlc:
	@echo "generating go code from sql"
	sqlc generate

## test: test the project
test:
	@echo "testing"
	go test -v -cover ./...

## server: start the HTTP server
server:
	@echo "starting the HTTP server"
	go run main.go

## mock: generates mock interfaces
mock:
	@echo "generating mock interfaces..."
	mockgen -package mockdb -destination db/mock/store.go github.com/foyez/simplebank/db/sqlc Store

.PHONY: db_docs db_schema postgres createdb dropdb create_migration migrateup migratedown sqlc test server mock