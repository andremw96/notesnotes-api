postgres:
	docker run --name postgres -p 5432:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=quipper123 -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root notesnotes

dropdb:
	docker exec -it postgres dropdb notesnotes

migrateup:
	migrate -path db/migration -database "postgresql://root:quipper123@localhost:5432/notesnotes?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:quipper123@localhost:5432/notesnotes?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown