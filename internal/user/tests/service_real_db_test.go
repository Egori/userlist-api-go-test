package user_test

import (
	"log"
	"os"
	"testing"

	user_pkg "userlist-api-test/internal/user"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testDB *gorm.DB

// Подключение к тестовой базе данных
func setupTestDB() {
	dsn := "host=localhost user=users_test password=654321 dbname=users_test port=5433 sslmode=disable"
	var err error
	testDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к тестовой базе данных: ", err)
	}
}

// Очистка базы данных перед каждым тестом
func clearTestDB() {
	testDB.Exec("TRUNCATE users RESTART IDENTITY CASCADE")
}

// Закрытие соединения с тестовой базой данных
func tearDownTestDB() {
	sqlDB, err := testDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}

func TestMain(m *testing.M) {
	setupTestDB()
	// Подготовка базы данных перед запуском тестов
	clearTestDB()

	// Запуск тестов
	code := m.Run()

	// Закрытие соединения с тестовой базой данных
	tearDownTestDB()

	os.Exit(code)
}

func TestCreateUserWithRealDB(t *testing.T) {
	// Инициализируем репозиторий с настоящей базой данных
	repo := user_pkg.NewRepository(testDB)
	service := user_pkg.NewService(repo)

	// Тестируем создание пользователя
	user := &user_pkg.User{Name: "John", Email: "john@example.com"}
	err := service.Create(user)

	// Проверяем, что ошибок нет
	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)

	// Проверяем, что пользователь добавлен в базу
	var savedUser user_pkg.User
	testDB.First(&savedUser, user.ID)
	assert.Equal(t, "John", savedUser.Name)
	assert.Equal(t, "john@example.com", savedUser.Email)
}

func TestCreateUserWithTransaction(t *testing.T) {
	// Начинаем транзакцию
	tx := testDB.Begin()
	if tx.Error != nil {
		t.Fatalf("Не удалось начать транзакцию: %v", tx.Error)
	}

	// Инициализируем репозиторий с транзакцией
	repo := user_pkg.NewRepository(tx)
	service := user_pkg.NewService(repo)

	// Тестируем создание пользователя
	user := &user_pkg.User{Name: "Jane", Email: "jane@example.com"}
	err := service.Create(user)
	assert.NoError(t, err)

	// Проверяем, что пользователь был сохранен
	var savedUser user_pkg.User
	tx.First(&savedUser, user.ID)
	assert.Equal(t, "Jane", savedUser.Name)

	// Откатываем транзакцию
	tx.Rollback()
}
