package tests

import (
	models "api/app/models/user"
	"api/app/router"
	"api/config"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/pressly/goose"
)

func MakeRequest(t *testing.T, request *http.Request) *httptest.ResponseRecorder {
	request.Header.Set("Content-Type", "application/json")

	router := router.GetRouter()

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response == nil {
		t.Fatal("Failed to execute request")
	}

	return response
}

func RunMigrations() {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println("error to connect database - ", err)
		panic(err)
	}

	file := filepath.Join(config.RootDir, "migrations/sqlite")

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, file); err != nil {
		log.Println("failed to run migrations: %w", err)
		panic(err)
	}
}

func CleanDb() error {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println("error to connect database - ", err)
	}

	file := filepath.Join(config.RootDir, "migrations/sqlite")

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Down(db, file); err != nil {
		return fmt.Errorf("error to clean db: %w", err)
	}

	return nil
}

func CreateUser() (models.UserModel, error) {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println("error to connect database - ", err)
	}

	statement, err := db.Prepare(
		"INSERT INTO users (name, nick, email, password) VALUES ('name', 'nick', 'email', 'password')",
	)
	if err != nil {
		return models.UserModel{}, err
	}
	defer statement.Close()

	result, err := statement.Exec()
	if err != nil {
		return models.UserModel{}, err
	}

	lastID, _ := result.LastInsertId()

	return FindUserById(uint64(lastID))
}

func FindUserById(ID uint64) (models.UserModel, error) {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println("error to connect database - ", err)
	}

	var user models.UserModel
	err = db.QueryRow(
		"SELECT id, name, nick, email, created_at FROM users WHERE id = ?",
		ID,
	).Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)

	if err != nil {
		return models.UserModel{}, err
	}

	return user, nil
}
