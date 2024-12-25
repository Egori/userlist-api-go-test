 
FROM alpine:latest

# Устанавливаем сертификаты
RUN apk add --no-cache ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем скомпилированный бинарный файл
COPY userlist-api .

RUN chmod +x userlist-api

# Копируем переменные окружения
COPY .env .env

# Открываем порт
EXPOSE 8080

# Запуск приложения
CMD ["/app/userlist-api"]
