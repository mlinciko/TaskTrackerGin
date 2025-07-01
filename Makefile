APP_NAME=task_tracker
DB_CONTAINER_NAME=postgres

.PHONY: run build test migrate docker-build docker-run docker-compose-build

# Server cmd commands (for local use)
run:
	set -a; source config/.env; set +a; go run cmd/main.go

build:
	mkdir -p ./bin
	go build -o ./bin/$(APP_NAME) ./cmd

test:
	go test -v ./...

migrate:
	set -a; source config/.env; set +a; go run cmd/migrate.go

# Run server with Dockerfile
docker-build:
	docker build -t $(APP_NAME) .

docker-run:
	docker run --env-file config/.env.prod -p 8080:8080 $(APP_NAME)

# Run server and DB with docker-compose
docker-compose-build:
	docker compose --env-file config/.env.prod build

docker-compose-up:
	docker compose --env-file config/.env.prod up
	
docker-compose-down:
	docker compose --env-file config/.env.prod down -v

