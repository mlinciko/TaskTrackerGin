# Стадия сборки
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o ./bin/task_tracker ./cmd

# Финальный образ
FROM alpine:latest

WORKDIR /app

# Устанавливаем корневые сертификаты (для подключения к БД через SSL, если нужно)
RUN apk --no-cache add ca-certificates

# Копируем собранный бинарник и .env файл
COPY --from=builder /app/bin/task_tracker .
COPY config/.env.prod ./config/.env

# Запуск приложения
CMD ["./task_tracker"]
