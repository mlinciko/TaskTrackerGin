APP_NAME=task_tracker
SOURCE_FOLDER=task_tracker
DB_CONTAINER_NAME=postgres

.PHONY: run build docker docker-run db-start db-stop db-clean migrate

run:
	set -a; source .env; set +a; go run $(SOURCE_FOLDER)/main.go

build:
	mkdir -p ./bin
	go build -o ./bin/$(APP_NAME) ./$(SOURCE_FOLDER)

docker-build:
	docker build -t $(APP_NAME) .

docker-run:
	docker run --env-file task_tracker/.env -p 8080:8080 $(APP_NAME)

db-start:
	docker run --name $(DB_CONTAINER_NAME) \
		--env-file .env \
		-p $$(grep DB_PORT .env | cut -d '=' -f2):5432 \
		-d postgres:15

db-stop:
	docker stop $(DB_CONTAINER_NAME) || true
	docker rm $(DB_CONTAINER_NAME) || true

db-clean: db-stop
	docker volume prune -f

migrate:
	set -a; source .env; set +a; go run cmd/migrate.go
