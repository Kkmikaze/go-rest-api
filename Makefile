run:
	go run cmd/main.go

mod:
	go mod download && go mod verify
	
migration:
	go run cmd/migration/migration.go

build:
	go build -o bin/main cmd/main.go

env:
	cp .env.example .env