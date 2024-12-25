package main

import (
	"fmt"
	"userlist-api-test/config"
	"userlist-api-test/internal/db"
	"userlist-api-test/internal/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Загрузка конфигурации
	config, err := config.LoadConfig()
	if err != nil {
		e.Logger.Fatal("ошибка при загрузке конфигурации: %v", err)
	}
	e.Logger.Debug(fmt.Sprintf("Loaded config: %+v", config))

	// Подключение к БД
	dbConn, err := db.ConnectDB(config)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Инициализация модулей
	var userRepo user.Repository = user.NewRepository(dbConn)
	var userService user.Service = user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:*"},                       // Разрешить запросы с любого порта localhost
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE}, // Разрешенные методы
	}))

	// Роутинг
	e.GET("/users", userHandler.GetAll)
	e.POST("/users", userHandler.Create)
	e.PUT("/users/:id", userHandler.Update)
	e.DELETE("/users/:id", userHandler.Delete)

	// Запуск сервера
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.HTTP_PORT)))
}
