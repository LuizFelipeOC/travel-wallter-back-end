package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/sijms/go-ora/v2"
)

func Connect() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	cnnStr := fmt.Sprintf("oracle://%s:%s@%s:%s/XE", dbUser, dbPassword, dbHost, dbPort)

	db, err := sql.Open("oracle", cnnStr)

	if err != nil {
		return nil, NewDatabaseError("DB_CONNECTION", "Erro ao conectar no banco", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)

	return db, nil
}