run:
	go run main.go

run migrate:
	go run db/migrate/migrate.go

build:
	go build -o bin/main main.go