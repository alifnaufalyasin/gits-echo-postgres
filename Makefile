migrateup:
	migrate -path migrations/ -database "postgresql://postgres:postgres@localhost:5432/gits-echo?sslmode=disable" up

migratedown:
	migrate -path migrations/ -database "postgresql://postgres:postgres@localhost:5432/gits-echo?sslmode=disable" -verbose down

postgres:
	docker-compose up

createdb:
	docker exec -it gits-echo-me_postgres_1 createdb --username=postgres --owner=postgres gits-echo

dropdb:
	docker exec -it gits-echo-me_postgres_1 dropdb --username=postgres gits-echo

run:
	go run server.go

.PHONY: migrateup migratedown postgres createdb dropdb run
