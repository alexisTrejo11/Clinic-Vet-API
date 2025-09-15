package db_test

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
	"unsafe"

	"clinic-vet-api/sqlc"
	"github.com/jackc/pgx/v5"
)

var testQueries *sqlc.Queries
var testDB *pgx.Conn

func InitTestDb() {
	var err error
	ctx := context.Background()

	testDB, err = pgx.Connect(ctx, "host=localhost port=5432 user=postgres password=postgrea dbname=clinic-vet_test sslmode=disable")
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
	InitTestDb()

	exitCode := m.Run()

	testDB.Close(context.Background())

	os.Exit(exitCode)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
