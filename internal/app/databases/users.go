package db

import (
	"database/sql"

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
	// user := os.Getenv("DATABASE_USER")
	// password := os.Getenv("DATABASE_PASSWORD")
	// dbname := os.Getenv("DATABASE_NAME")
	// host := os.Getenv("DATABASE_HOST")
	// port := os.Getenv("DATABASE_PORT")
	// connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
	// 	user, password, dbname, host, port)
	db, err := sql.Open("postgres", "user=bobhuynh password=postgres dbname=postgres host=192.168.215.2 port=5432 sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &PostgreSQLDB{db}, nil
}
