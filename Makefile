postgres:
	docker run --name postgres16 --network fitnessapp-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root fitness_app

dropdb:
	docker exec -it postgres16 dropdb fitness_app

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fitness_app?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fitness_app?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	sudo go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server