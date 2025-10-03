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

func PtrToCustomerIDPtr(id *uint) *valueobject.CustomerID {
	if id == nil {
		return nil
	}
	customerID := valueobject.NewCustomerID(*id)
	return &customerID
}

func PtrToMedSessionIDPtr(id *uint) *valueobject.MedSessionID {
	if id == nil {
		return nil
	}
	medSessionID := valueobject.NewMedSessionID(*id)
	return &medSessionID
}

func MedSessionIDPtrToPtr(id *valueobject.MedSessionID) *uint {
	if id == nil {
		return nil
	}
	medSessionID := id.Value()
	return &medSessionID
}

func PaidByCustomerPtrToPtr(id *valueobject.CustomerID) *uint {
	if id == nil {
		return nil
	}
	customerID := id.Value()
	return &customerID
}

func PaidByEmployeePtrToPtr(id *valueobject.EmployeeID) *uint {
	if id == nil {
		return nil
	}
	employeeID := id.Value()
	return &employeeID
}

func MoneyPtrToFloat64Ptr(money *valueobject.Money) *float64 {
	if money == nil {
		return nil
	}
	amount := money.Amount().Float64()
	return &amount
}

func PaidToEmployeePtrToPtr(id *valueobject.EmployeeID) *uint {
	if id == nil {
		return nil
	}
	employeeID := id.Value()
	return &employeeID
}
