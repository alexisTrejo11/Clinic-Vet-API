createTestDB:
	migrate -path db/migration -database "postgresql://postgres:root@localhost:5430/vet_database_test?sslmode=disable" -verbose up