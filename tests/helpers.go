package tests

import (
	"api/app/router"
	"api/config"
	"database/sql"
	"fmt"
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

func RunMigrations(db *sql.DB) error {
	file := filepath.Join(config.RootDir, "migrations/sqlite")

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, file); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

func CleanDb(db *sql.DB) error {
	file := filepath.Join(config.RootDir, "migrations/sqlite")

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Down(db, file); err != nil {
		return fmt.Errorf("error to clean db: %w", err)
	}

	return nil
}
