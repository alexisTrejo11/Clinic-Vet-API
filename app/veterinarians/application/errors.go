package vetApplication

import (
	appError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
	infra_error "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/infrastructure"
)

func VetNotFoundErr(parameter string, value string) error {
	resource := "veterinarian " + parameter
	return appError.NewConflictError(resource, value)
}

func VetDBErr(operation string, err error) error {
	return infra_error.NewDatabaseError(operation, err.Error())
}
