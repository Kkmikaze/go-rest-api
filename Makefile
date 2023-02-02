run:
	go run main.go
	
migrate:
	go run db/migrate/migrate.go

build:
	go build -o bin/main main.go

env:
	cp .env.example .env