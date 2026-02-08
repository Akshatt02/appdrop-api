// Package db manages PostgreSQL database connections and pooling.
// It provides a connection pool for efficient database operations throughout the application.
package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Pool is the global database connection pool.
// It handles concurrent database requests efficiently and should be used
// across the entire application for all database operations.
var Pool *pgxpool.Pool

// ConnectDB initializes the PostgreSQL connection pool from the DATABASE_URL environment variable.
// It creates a new connection pool, verifies connectivity with a ping,
// and makes it available globally via the Pool variable.
// Panics if connection fails or DATABASE_URL is not set.
func ConnectDB() {
	dbURL := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Postgres")
	Pool = pool
}
