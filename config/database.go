package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

func ConnectMysql() (*sql.DB, error) {
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
