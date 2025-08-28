package vetApplication

import (
	ApplicationError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
	infraErr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/infrastructure"
)

func VetNotFoundErr(parameter string, value string) error {
	resource := "veterinarian " + parameter
	return ApplicationError.NewConflictError(resource, value)
}

func VetDBErr(operation string, err error) error {
	return infraErr.NewDatabaseError(operation, err.Error())
}
