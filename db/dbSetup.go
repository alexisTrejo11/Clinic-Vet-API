package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

var db *pgx.Conn

func InitDb() *pgx.Conn {
	var err error

	// Read environment variables for database connection details
	dbHost := "postgres12"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "root"
	dbName := "vet_database"

	// Establish connection to the database
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err = pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// Check if the connection is successful
	if err := db.Ping(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Error pinging database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to the database successfully")

	return db
}
