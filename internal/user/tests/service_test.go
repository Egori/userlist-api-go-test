package user_test

import (
	"testing"
	user_pkg "userlist-api-test/internal/user"

	"github.com/stretchr/testify/assert"
)

// mockRepository is a mock implementation of the repository interface for testing purposes.
func TestCreateUser(t *testing.T) {
	// Создаем мок-репозиторий
	mockRepo := &mockRepository{}
	// Создаем сервис, передавая мок-репозиторий
	service := user_pkg.NewService(mockRepo)

	// Тестируем создание пользователя
	user := &user_pkg.User{Name: "John", Email: "john@example.com"}
	err := service.Create(user)

	// Проверяем, что ошибок нет и ID пользователя установлен
	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
}

func TestCreateUserWithEmptyName(t *testing.T) {
	// Создаем мок-репозиторий
	mockRepo := &mockRepository{}
	// Создаем сервис
	service := user_pkg.NewService(mockRepo)

	// Тестируем создание пользователя с пустым именем
	user := &user_pkg.User{Name: "", Email: "john@example.com"}
	err := service.Create(user)

	// Проверяем, что ошибка произошла
	assert.Error(t, err)
	assert.Equal(t, "Name cannot be empty", err.Error())
}

func TestGetAllUsers(t *testing.T) {
	// Создаем мок-репозиторий
	mockRepo := &mockRepository{}
	// Создаем сервис
	service := user_pkg.NewService(mockRepo)

	// Тестируем получение всех пользователей
	users, err := service.GetAll()

	// Проверяем, что ошибок нет и данные возвращены правильно
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "John", users[0].Name)
}
