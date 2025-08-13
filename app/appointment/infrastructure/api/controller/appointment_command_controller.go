package appointmentController

import (
	"context"
	"errors"
	"net/http"

	appointmentCmd "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	responses "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AppointmentCommandController struct {
	commandBus appointmentCmd.CommandBus
	validate   *validator.Validate
}

func NewAppointmentCommandController(
	commandBus appointmentCmd.CommandBus,
	validate *validator.Validate,
) *AppointmentCommandController {
	return &AppointmentCommandController{
		commandBus: commandBus,
		validate:   validate,
	}
}

// Create appointment
func (controller *AppointmentCommandController) CreateAppointment(ctx *gin.Context) {
	var command appointmentCmd.CreateAppointmentCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		responses.RequestBodyDataError(ctx, err)
		return
	}

	result := controller.commandBus.Execute(ctx, command)
	if !result.IsSuccess {
		responses.ApplicationError(ctx, result.Error)
		return
	}

	responses.Created(ctx, gin.H{"message": result.Message, "appointment_id": result.Id})
}

// Update appointment
func (controller *AppointmentCommandController) UpdateAppointment(ctx *gin.Context) {
	idInt, err := shared.ParseID(ctx, "id")
	if err != nil {
		responses.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	var command appointmentCmd.UpdateAppointmentCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.validate.Struct(command); err != nil {
		responses.RequestBodyDataError(ctx, err)
		return
	}

	command.AppointmentId = idInt
	result := controller.commandBus.Execute(ctx, command)
	if !result.IsSuccess {
		responses.ApplicationError(ctx, result.Error)
		return
	}

	responses.Success(ctx, result.Message)
}

// Delete appointment
func (controller *AppointmentCommandController) DeleteAppointment(ctx *gin.Context) {
	idInt, err := shared.ParseID(ctx, "id")
	if err != nil {
		responses.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	command := appointmentCmd.NewDeleteAppointmentCommand(idInt)

	result := controller.commandBus.Execute(ctx, command)
	if !result.IsSuccess {
		responses.ApplicationError(ctx, result.Error)
		return
	}

	responses.Success(ctx, result.Message)
}

// Reschedule appointment
func (controller *AppointmentCommandController) RescheduleAppointment(ctx *gin.Context) {
	appointmentId, err := shared.ParseID(ctx, "id")
	if err != nil {
		responses.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	var command appointmentCmd.RescheduleAppointmentCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		responses.RequestBodyDataError(ctx, err)
		return
	}

	command.AppointmentId = appointmentId

	result := controller.commandBus.Execute(context.Background(), command)
	if !result.IsSuccess {
		responses.ApplicationError(ctx, result.Error)
		return
	}

	responses.Success(ctx, result.Message)
}

// MarkAsNoShow allows veterinarians to mark appointments as no-show
func (controller *AppointmentCommandController) MarkAsNoShow(ctx *gin.Context) {
	appointmentId, err := shared.ParseID(ctx, "id")
	if err != nil {
		responses.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	command := appointmentCmd.MarkAsNotPresentedCommand{
		Id: appointmentId,
	}

	result := controller.commandBus.Execute(context.Background(), command)
	if !result.IsSuccess {
		responses.ApplicationError(ctx, result.Error)
		return
	}

	responses.Success(ctx, result)
}

func (controller *AppointmentCommandController) ConfirmAppointment(ctx *gin.Context) {
	appointmentId, err := shared.ParseID(ctx, "id")
	if err != nil {
		responses.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	// Get vet id from context
	vetIdInterface, exists := ctx.Get("vet_id")
	if !exists {
		responses.Unauthorized(ctx, errors.New("vet id not found in context"))
		return
	}

	vetIdInt, ok := vetIdInterface.(int)
	if !ok {
		responses.BadRequest(ctx, errors.New("invalid vet id format"))
		return
	}

	vetId, err := vetDomain.NewVeterinarianId(vetIdInt)
	if err != nil {
		responses.BadRequest(ctx, err)
		return
	}

	command := appointmentCmd.ConfirmAppointmentCommand{
		Id:    appointmentId,
		VetId: &vetId,
	}

	result := controller.commandBus.Execute(context.Background(), command)
	if !result.IsSuccess {
		responses.ApplicationError(ctx, result.Error)
		return
	}

	responses.Success(ctx, result)
}

// CompleteAppointment allows veterinarians to mark appointments as completed
// PUT /vet/appointments/:id/complete
func (controller *AppointmentCommandController) CompleteAppointment(ctx *gin.Context) {
	appointmentId, err := shared.ParseID(ctx, "id")
	if err != nil {
		responses.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	var requestBody struct {
		Notes *string `json:"notes,omitempty"`
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		responses.RequestBodyDataError(ctx, err)
		return
	}

	command := appointmentCmd.CompleteAppointmentCommand{
		Id:    appointmentId,
		Notes: requestBody.Notes,
	}

	result := controller.commandBus.Execute(context.Background(), command)
	if !result.IsSuccess {
		responses.ApplicationError(ctx, result.Error)
		return
	}

	responses.Success(ctx, result)
}
