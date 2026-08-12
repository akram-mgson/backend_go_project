package main

import (
	"database/sql"
	"errors"
	"os"

	// go get -u github.com/go-sql-driver/mysql - я скачал драйвер

	// обязательно импортирование драйверов
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {

	dsn := os.Getenv("DB_DSN")

	if dsn == "" {
		return nil, errors.New("ДБ окружение еще не установлено")
	}

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil

}
