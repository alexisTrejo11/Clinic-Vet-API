package mapper

import "clinic-vet-api/app/modules/core/domain/valueobject"

type ValueObjcetMapper struct {
}

func PtrToEmployeeIDPtr(id *uint) *valueobject.EmployeeID {
	if id == nil {
		return nil
	}
	employeeID := valueobject.NewEmployeeID(*id)
	return &employeeID
}
