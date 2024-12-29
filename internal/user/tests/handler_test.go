package user_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	user_pkg "userlist-api-test/internal/user"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}
	db.AutoMigrate(&user_pkg.User{})
	// Очищаем таблицу перед каждым тестом
	db.Exec("DELETE FROM users")
	return db
}

func setupHandler(db *gorm.DB) *user_pkg.Handler {
	var userRepo user_pkg.Repository = user_pkg.NewRepository(db)
	var userService user_pkg.Service = user_pkg.NewService(userRepo)
	return user_pkg.NewHandler(userService)
}

func TestGetAllUsers(t *testing.T) {
	db := setupTestDB()
	handler := setupHandler(db)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.GetAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "[]\n", rec.Body.String())
	}
}

func TestGetAllUsersWithData(t *testing.T) {
	db := setupTestDB()
	handler := setupHandler(db)

	// Создаем тестовых пользователей
	users := []user_pkg.User{
		{Name: "John Doe", Email: "john@example.com"},
		{Name: "Jane Doe", Email: "jane@example.com"},
	}
	for _, u := range users {
		db.Create(&u)
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Вызываем эндпоинт
	if assert.NoError(t, handler.GetAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// Декодируем ответ
		var response []user_pkg.User
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		// Проверяем, что количество пользователей совпадает
		assert.Equal(t, len(users), len(response))

		// Проверяем, что каждый пользователь в ответе соответствует ожидаемому
		for i, expectedUser := range users {
			assert.Equal(t, expectedUser.Name, response[i].Name)
			assert.Equal(t, expectedUser.Email, response[i].Email)
			assert.NotZero(t, response[i].ID)
			assert.NotZero(t, response[i].CreatedAt)
			assert.NotZero(t, response[i].UpdatedAt)
		}
	}
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB()
	handler := setupHandler(db)

	e := echo.New()
	user := user_pkg.User{Name: "John Doe", Email: "john@example.com"}
	userJSON, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var response user_pkg.User
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, user.Name, response.Name)
		assert.Equal(t, user.Email, response.Email)
	}
}

func TestUpdateUser(t *testing.T) {
	db := setupTestDB()
	handler := setupHandler(db)

	// Создаем пользователя для обновления
	user := user_pkg.User{Name: "John Doe", Email: "john@example.com"}
	db.Create(&user)

	e := echo.New()
	updatedUser := user_pkg.User{Name: "Jane Doe", Email: "jane@example.com"}
	updatedUserJSON, _ := json.Marshal(updatedUser)
	req := httptest.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(updatedUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, handler.Update(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response user_pkg.User
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, updatedUser.Name, response.Name)
		assert.Equal(t, updatedUser.Email, response.Email)
	}
}

func TestDeleteUser(t *testing.T) {
	db := setupTestDB()
	handler := setupHandler(db)

	// Создаем пользователя для удаления
	user := user_pkg.User{Name: "John Doe", Email: "john@example.com"}
	db.Create(&user)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, handler.Delete(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}
