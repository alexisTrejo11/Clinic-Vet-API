TablesMigrationUp:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/clinic-vet?sslmode=disable" -verbose up

TablesMigrationDown:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/clinic-vet?sslmode=disable" -verbose down

CreateTablesMigrationTest:
	migrate -path db/migrations/test/ -database "postgresql://postgres:root@localhost:5432/clinic-vet_test?sslmode=disable" -verbose up

InitSwagger:
	swag init

SQLC:
	sqlc init

CREATESQLCMOCK:
	mockgen -source=sqlc/querier.go -destination=test/mock/queries_mock.go -package=mock