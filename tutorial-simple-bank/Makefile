MIGRATE_NAME?=update_table
DB_URL?=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name simple-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it simple-postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it simple-postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

create-new-migration:
	migrate create -ext sql -dir db/migration -seq $(MIGRATE_NAME)

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock-store:
	mockgen -package mockdb -destination db/mock/store.go github.com/tutorial/simple-bank/db/sqlc Store 

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock-store proto redis