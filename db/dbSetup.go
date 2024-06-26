package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

var db *pgx.Conn

func InitDb() {
	var err error

	// Establish connection to the database
	db, err = pgx.Connect(context.Background(), "host=localhost port=5431 user=postgres dbname=vet_database password=vladilena sslmode=disable")
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
}
