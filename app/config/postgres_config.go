package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseConfig struct {
	URL             string        `json:"url"`
	MaxOpenConns    int           `json:"max_open_conns"`
	MaxIdleConns    int           `json:"max_idle_conns"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `json:"conn_max_idle_time"`
}

func PostgresConn(dbURL string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func CreatePgxPool(dbURL string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse database config: %v\n", err)
		os.Exit(1)
	}

	// Configurar el pool de conexiones
	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnLifetime = 1 * time.Hour
	config.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	// Verificar la conexión
	if err := pool.Ping(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully connected to PostgreSQL database with pgx pool!")
	return pool
}

func loadDatabaseConfig(config *DatabaseConfig) error {
	config.URL = os.Getenv("DATABASE_URL")
	if config.URL == "" {
		return fmt.Errorf("DATABASE_URL is required")
	}

	var err error
	config.MaxOpenConns, err = parseIntWithDefault("DB_MAX_OPEN_CONNS", 25)
	if err != nil {
		return fmt.Errorf("invalid DB_MAX_OPEN_CONNS: %w", err)
	}

	config.MaxIdleConns, err = parseIntWithDefault("DB_MAX_IDLE_CONNS", 5)
	if err != nil {
		return fmt.Errorf("invalid DB_MAX_IDLE_CONNS: %w", err)
	}

	config.ConnMaxLifetime, err = parseDuration("DB_CONN_MAX_LIFETIME", "1h")
	if err != nil {
		return fmt.Errorf("invalid DB_CONN_MAX_LIFETIME: %w", err)
	}

	config.ConnMaxIdleTime, err = parseDuration("DB_CONN_MAX_IDLE_TIME", "30m")
	if err != nil {
		return fmt.Errorf("invalid DB_CONN_MAX_IDLE_TIME: %w", err)
	}

	return nil
}
