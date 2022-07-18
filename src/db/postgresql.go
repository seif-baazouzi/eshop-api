package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type connection struct {
	db        *sql.DB
	isRunning bool
}

var connections [10]connection

func InitPostgresql() {
	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	for i := 0; i < 10; i++ {
		db, err := sql.Open("postgres", databaseUrl)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}

		connections[i].db = db
		connections[i].isRunning = false
	}
}

func GetPool() *sql.DB {
	for _, connection := range connections {
		if !connection.isRunning {
			return connection.db
		}
	}

	time.Sleep(5_000_000) // 5ms
	return GetPool()
}

func ClosePool(db *sql.DB) {
	for _, connection := range connections {
		if connection.db == db {
			connection.isRunning = false
			break
		}
	}
}

func ClosePostgresql() {
	for _, connection := range connections {
		connection.db.Close()
	}
}
