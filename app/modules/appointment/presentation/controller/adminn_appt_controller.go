package controller

import (
	"clinic-vet-api/app/modules/appointment/infrastructure/bus"
	"clinic-vet-api/app/shared/error/infrastructure/http"
	ginutils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminApptController struct {
	bus        *bus.AppointmentBus
	validator  *validator.Validate
	operations *ApptControllerOperations
}

func NewAdminApptController(
	bus *bus.AppointmentBus,
	validator *validator.Validate,
	operations *ApptControllerOperations,
) *AdminApptController {
	return &AdminApptController{
		bus:        bus,
		validator:  validator,
		operations: operations,
	}
}

func (ctrl *AdminApptController) GetBySpecfificationAppointments(c *gin.Context) {

}
func (ctrl *AdminApptController) GetAppointmentByID(c *gin.Context) {
	ctrl.operations.FindAppointmentByID(c, GetByIDExtraArgs{})
}
func (ctrl *AdminApptController) CreateAppointment(c *gin.Context) {
	employeeID := c.Param("employeeID")

	if employeeID == "" {
		response.BadRequest(c, http.RequestURLQueryError(errors.New("employeeID query param is required"), c.Request.URL.RawQuery))
		return
	}

	employeeIDUint, err := strconv.ParseUint(employeeID, 10, 64)
	if err != nil {
		response.BadRequest(c, http.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	employeeIDUintVal := uint(employeeIDUint)
	ctrl.operations.CreateAppointment(c, &employeeIDUintVal)
}
func (ctrl *AdminApptController) UpdateAppointment(c *gin.Context) {
	appointmentID, err := ginutils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, http.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}
	ctrl.operations.UpdateAppointment(c, appointmentID)
}
func (ctrl *AdminApptController) DeleteAppointment(c *gin.Context) {
	ctrl.operations.DeleteAppointment(c)
}
