package mapper

import "clinic-vet-api/app/modules/core/domain/valueobject"

type ValueObjectMapper struct{}

func (m *ValueObjectMapper) PaymentIDToUint(id valueobject.PaymentID) uint {
	return id.Value()
}

func (m *ValueObjectMapper) UintToPaymentID(id uint) valueobject.PaymentID {
	return valueobject.NewPaymentID(id)
}

func (m *ValueObjectMapper) CustomerIDToUint(id valueobject.CustomerID) uint {
	return id.Value()
}

func (m *ValueObjectMapper) UintToCustomerID(id uint) valueobject.CustomerID {
	return valueobject.NewCustomerID(id)
}

func (m *ValueObjectMapper) EmployeeIDToUint(id valueobject.EmployeeID) uint {
	return id.Value()
}

func (m *ValueObjectMapper) MoneyPtrToFloat64Ptr(money *valueobject.Money) *float64 {
	if money != nil {
		amount := money.Amount().Float64()
		return &amount
	}
	return nil
}
