package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbHost := getEnv("DB_HOST", "postgres12")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "vet_database")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "root")

	databaseURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	fmt.Printf("Running migrations on: %s\n", databaseURL)

	// Wait for database to be ready
	if err := waitForDatabase(databaseURL, 30*time.Second); err != nil {
		log.Fatalf("❌ Database not ready: %v", err)
	}

	// Run migrations
	if err := runMigrations(databaseURL); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	fmt.Println("Migrations completed successfully!")
}

func runMigrations(databaseURL string) error {
	m, err := migrate.New(
		"file://db/migrations",
		databaseURL,
	)
	if err != nil {
		return fmt.Errorf("create migrate instance: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrate up: %w", err)
	}

	return nil
}

func waitForDatabase(databaseURL string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for database")
		default:
			pool, err := pgxpool.New(ctx, databaseURL)
			if err == nil {
				defer pool.Close()
				if err := pool.Ping(ctx); err == nil {
					return nil
				}
			}
			time.Sleep(2 * time.Second)
		}
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
