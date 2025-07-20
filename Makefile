.PHONY:
	create_migration migrate migrate_local generate

generate:
	sqlc generate

default:
	go vet ./...

migrate:
	migrate -database $(database) -source file://./db/migrations/ up

migrate_local:
	migrate -database "postgres://postgres@localhost:5432/postgres?sslmode=disable" -source file://./db/migrations/ up

create_migration:
	migrate create -ext sql -dir db/migrations -seq -digits 3 $(name)

