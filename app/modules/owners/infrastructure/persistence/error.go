package persistence

import (
	infraErr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/infrastructure"
)

func DBCreateError(message string) *infraErr.DatabaseError {
	return infraErr.NewDatabaseError("create owner enitty", message)
}

func DBDeleteError(message string) *infraErr.DatabaseError {
	return infraErr.NewDatabaseError("delete owner enitty", message)
}

func DBUpdateError(message string) *infraErr.DatabaseError {
	return infraErr.NewDatabaseError("update owner enitty", message)
}

func DBSelectFoundError(message string) *infraErr.DatabaseError {
	return infraErr.NewDatabaseError("delete owner enitty", message)
}
