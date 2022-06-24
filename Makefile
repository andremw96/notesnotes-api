postgres:
	docker run --name postgres --network notesnotes-network -p 5432:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=quipper123 -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root notesnotes

dropdb:
	docker exec -it postgres dropdb notesnotes

createnewmigration:
	migrate create -ext sql -dir db/migration -seq $(ARGS)

migrateup:
	migrate -path db/migration -database "postgresql://root:quipper123@localhost:5432/notesnotes?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:quipper123@localhost:5432/notesnotes?sslmode=disable" -verbose down

sqlc:
	sqlc generate

unit_test_run:
	go test -v -coverpkg=./... -coverprofile=coverage.out ./...

startserver:
	go run main.go

generatemock:
	mockgen -package mockdb -destination db/mock/store.go andre/notesnotes-api/db/sqlc Store

createdockerimage:
	docker build -t notesnotes-api:latest .

rundockerimagedebug:
	docker run --name notesnotes-api -p 8080:8080 notesnotes-api:latest

rundockerimagerelease:
	docker run --name notesnotes-api --network notesnotes-network -p 8080:8080 -e DB_SOURCE=postgresql://root:quipper123@postgres:5432/notesnotes?sslmode=disable -e GIN_MODE=release notesnotes-api:latest

removedockercontainer:
	docker rm notesnotes-api

.PHONY: postgres createdb dropdb createnewmigration migrateup migratedown sqlc unit_test_run startserver generatemock createdockerimage rundockerimagedebug rundockerimagerelease removedockercontainer