package persistence

import dberr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/database"

func DBCreateError(message string) error {
	return dberr.DatabaseOperationError("create", "pets", "sql", "failed to create record: "+message)
}

func DBDeleteError(message string) error {
	return dberr.DatabaseOperationError("delete", "pets", "sql", "failed to delete record: "+message)
}

func DBUpdateError(message string) error {
	return dberr.DatabaseOperationError("update", "pets", "sql", "failed to update record: "+message)
}

func DBSelectFoundError(message string) error {
	return dberr.DatabaseOperationError("select", "pets", "sql", "failed to find record: "+message)
}
