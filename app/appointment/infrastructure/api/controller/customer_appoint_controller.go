package controller

type CustomerAppointmentController struct {
}

func NewCustomerAppointmentController() *CustomerAppointmentController {
	return &CustomerAppointmentController{}
}

func (c *CustomerAppointmentController) GetMyAppointment() error {
	return nil
}

func (c *CustomerAppointmentController) CreateAppointment() error {
	return nil
}

func (c *CustomerAppointmentController) ConfirmMyAppointment() error {
	return nil
}

func (c *CustomerAppointmentController) CancelMyAppointment() error {
	return nil
}
