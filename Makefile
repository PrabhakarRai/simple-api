postgres:
	docker run --name postgres13 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root simple_api

dropdb:
	docker exec -it postgres13 dropdb simple_api

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_api?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_api?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_api?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_api?sslmode=disable" -verbose down 1

psql:
	docker exec -it postgres13 psql -U root

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres psql createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server