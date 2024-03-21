package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Database interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Close() error
}

type PostgreSQLDB struct {
	*sql.DB
}

func NewPostgreSQLDB() (Database, error) {
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	fmt.Println("DATABASE_USER:", dbUser)
	fmt.Println("DB_PASSWORD:", dbPassword)
	fmt.Println("DATABASE_NAME:", dbName)
	fmt.Println("DATABASE_HOST:", dbHost)
	fmt.Println("DATABASE_PORT:", dbPort)
	fmt.Println("DB_SSL_MODE:", dbSSLMode)

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort, dbSSLMode)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return &PostgreSQLDB{db}, nil
}
