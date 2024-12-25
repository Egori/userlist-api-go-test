# Build Stage
FROM golang:1.23.2-alpine as builder

WORKDIR /app

# Копируем только файлы модулей для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod tidy && go mod download

# Копируем остальной код
COPY . .

# Сборка бинарного файла
RUN go build -o userlist-api ./cmd/app

# Runtime Stage
FROM alpine:latest

# Устанавливаем сертификаты
RUN apk add --no-cache ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем только бинарный файл
COPY --from=builder /app/userlist-api .

COPY .env .env


# Открываем порт
EXPOSE 8080

# Запуск приложения
CMD ["/app/userlist-api"]
