package db_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5"
)

var testQueries *sqlc.Queries
var testDB *pgx.Conn

func InitTestDb() {
	var err error
	ctx := context.Background()

	testDB, err = pgx.Connect(ctx, "host=localhost port=5430 user=postgres password=root dbname=vet_database_test sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to test database: %v\n", err)
		os.Exit(1)
	}

	if err := testDB.Ping(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Error pinging test database: %v\n", err)
		os.Exit(1)
	}

	testQueries = sqlc.New(testDB)
}

func TestMain(m *testing.M) {
	// Initialize the test database
	InitTestDb()

	// Run tests
	exitCode := m.Run()

	// Close database connection after tests
	testDB.Close(context.Background())

	os.Exit(exitCode)
}
