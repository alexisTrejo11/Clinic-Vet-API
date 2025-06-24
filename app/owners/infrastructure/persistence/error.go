package sqlcOwnerRepository

import (
	infra_error "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/infrastructure"
)

func DBCreateError(message string) *infra_error.DatabaseError {
	return infra_error.NewDatabaseError("create owner enitty", message)
}

func DBDeleteError(message string) *infra_error.DatabaseError {
	return infra_error.NewDatabaseError("delete owner enitty", message)
}

func DBUpdateError(message string) *infra_error.DatabaseError {
	return infra_error.NewDatabaseError("update owner enitty", message)
}

func DBSelectFoundError(message string) *infra_error.DatabaseError {
	return infra_error.NewDatabaseError("delete owner enitty", message)
}
