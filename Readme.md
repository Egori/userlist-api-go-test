# REST API для работы с пользователями

## Описание
Это приложение представляет собой REST API для управления пользователями. Оно реализовано на языке Go с использованием фреймворка Echo и ORM библиотеки Gorm. API предоставляет следующие эндпоинты:

- **GET /users** - Получение списка пользователей.
- **POST /users** - Добавление нового пользователя.
- **PUT /users/:id** - Обновление информации о пользователе.
- **DELETE /users/:id** - Удаление пользователя.

Приложение работает в Docker-контейнере, с использованием PostgreSQL в качестве базы данных.

## Как запустить проект

### Шаг 1. Запуск Docker Compose
Убедитесь, что Docker и Docker Compose установлены на вашем компьютере. Запустите приложение с помощью команды:
```bash
docker-compose up --build
```

### Шаг 2. Проверка работы API
После успешного запуска, API будет доступен по адресу: `http://localhost:8080`. Вы можете использовать Postman, curl или любой другой инструмент для тестирования эндпоинтов.

### Шаг 3. Выполнение тестов
Для запуска тестов выполните следующую команду внутри контейнера приложения:
```bash
docker exec -it userlist_test_app go test ./...
```

## Инфраструктура
Проект использует многоконтейнерную инфраструктуру с помощью Docker Compose:

- **PostgreSQL**: используется как база данных, контейнер `userlist_test_pg` . 
- **Приложение на Go**: обрабатывает HTTP-запросы и взаимодействует с базой данных через Gorm, контейнер `userlist_test_app`.


## Примечания
- Убедитесь, что порт 5433 не занят другим процессом на вашем компьютере.
- Если нужно очистить контейнеры, выполните команду:
  ```bash
  docker-compose down -v
  ```

