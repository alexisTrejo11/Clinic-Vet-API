package command

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type RequestApptByCustomerCommand struct {
	petID         valueobject.PetID
	customerID    valueobject.CustomerID
	requestedDate time.Time
	service       enum.ClinicService
	notes         *string
}

func NewRequestApptByCustomerCommand(
	petID uint, customerID uint, requestedDate time.Time, service string, notes *string,
) (RequestApptByCustomerCommand, error) {
	cmd := &RequestApptByCustomerCommand{
		petID:         valueobject.NewPetID(petID),
		requestedDate: requestedDate,
		customerID:    valueobject.NewCustomerID(customerID),
		service:       enum.ClinicService(service),
		notes:         notes,
	}

	if err := cmd.validate(); err != nil {
		return RequestApptByCustomerCommand{}, err
	}

	return *cmd, nil
}

func (c *RequestApptByCustomerCommand) validate() error {
	if c.petID.IsZero() {
		return requestScheduleCmdErr("pet_id", "Pet ID is required")
	}
	if c.customerID.IsZero() {
		return requestScheduleCmdErr("customer_id", "Customer ID is required")
	}
	if c.requestedDate.IsZero() {
		return requestScheduleCmdErr("requested_date", "Requested date is required")
	}
	if !c.service.IsValid() {
		return requestScheduleCmdErr("service", "Service is invalid")
	}
	return nil
}

func (c *RequestApptByCustomerCommand) PetID() valueobject.PetID           { return c.petID }
func (c *RequestApptByCustomerCommand) CustomerID() valueobject.CustomerID { return c.customerID }
func (c *RequestApptByCustomerCommand) RequestedDate() time.Time           { return c.requestedDate }
func (c *RequestApptByCustomerCommand) Service() enum.ClinicService        { return c.service }
func (c *RequestApptByCustomerCommand) Notes() *string                     { return c.notes }
