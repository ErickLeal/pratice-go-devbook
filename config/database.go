package config

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectDatabase() (*sql.DB, error) {
	switch DBConnection {
	case "mysql":
		return connectMysql()
	case "sqlite":
		return connectSQLite()
	default:
		return nil, fmt.Errorf("database connection for: %s, is not configured", DBConnection)
	}
}

func connectMysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", DBConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}

func connectSQLite() (*sql.DB, error) {
	projectRoot := RootDir

	dbPath := filepath.Join(projectRoot, "tests/devbook.db")
	dbDir := filepath.Dir(dbPath)

	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		err := os.MkdirAll(dbDir, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("failed to create sqlite data file: %w", err)
		}
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite connection: %w", err)
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, fmt.Errorf("error to enable foreign keys: %w", err)
	}

	return db, nil
}
