APP_NAME=task_tracker
DB_CONTAINER_NAME=postgres

.PHONY: run build docker docker-run db-start db-stop db-clean migrate

# Server cmd commands (for local use)
run:
	set -a; source config/.env.dev; set +a; go run cmd/main.go

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
	docker run --env-file config/.env -p 8080:8080 $(APP_NAME)

# Run server and DB with docker-compose
# db-start:
# 	docker run --name $(DB_CONTAINER_NAME) \
# 		--env-file .env \
# 		-p $$(grep DB_PORT .env | cut -d '=' -f2):5432 \
# 		-d postgres:15

# db-stop:
# 	docker stop $(DB_CONTAINER_NAME) || true
# 	docker rm $(DB_CONTAINER_NAME) || true

# db-clean: db-stop
# 	docker volume prune -f

