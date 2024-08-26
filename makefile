TablesMigrationUp:
	migrate -path db/migration -database "postgresql://postgres:root@localhost:5431/vet_database?sslmode=disable" -verbose up

TablesMigrationDown:
	migrate -path db/migration -database "postgresql://postgres:root@localhost:5431/vet_database?sslmode=disable" -verbose down


CreateTablesMigrationTest:
	migrate -path db/migration/test/ -database "postgresql://postgres:root@localhost:5430/vet_database_test?sslmode=disable" -verbose up

InitSwagger:
	swag init

SQLC:
	sqlc init