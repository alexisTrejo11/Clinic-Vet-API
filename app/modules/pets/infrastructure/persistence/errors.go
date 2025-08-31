package persistence

import (
	infraErr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/infrastructure"
)

func DBCreateError(message string) *infraErr.DatabaseError {
	return infraErr.NewDatabaseError("create", message)
}

func DBDeleteError(message string) *infraErr.DatabaseError {
	return infraErr.NewDatabaseError("delete", message)
}

func DBUpdateError(message string) *infraErr.DatabaseError {
	return infraErr.NewDatabaseError("update", message)
}

func DBSelectFoundError(message string) *infraErr.DatabaseError {
	return infraErr.NewDatabaseError("delete", message)
}
