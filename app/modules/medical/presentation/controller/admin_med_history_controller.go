// Package controller defines the controllers for handling HTTP requests related to medical histories.
package controller

import (
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/modules/medical/application/command"
	"clinic-vet-api/app/modules/medical/application/query"
	"clinic-vet-api/app/modules/medical/infrastructure/bus"
	"clinic-vet-api/app/modules/medical/presentation/dto"
	"clinic-vet-api/app/shared/response"
	"errors"

	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminMedicalHistoryController struct {
	bus       *bus.MedicalHistoryBus
	validator *validator.Validate
}

func NewAdminMedicalHistoryController(bus *bus.MedicalHistoryBus) *AdminMedicalHistoryController {
	return &AdminMedicalHistoryController{
		bus:       bus,
		validator: validator.New(),
	}
}

func (ctlr AdminMedicalHistoryController) SearchMedicalHistories(c *gin.Context) {
}

func (ctlr AdminMedicalHistoryController) GetMedicalHistoryDetails(c *gin.Context) {
	idUint, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "medical-history", c.Param("id")))
		return
	}

	query := query.NewFindMedHistByIDQuery(idUint, c.Request.Context())

	medHistory, err := ctlr.bus.QueryBus.FindMedHistByID(*query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Found(c, medHistory, "Medical History")
}

func (ctlr AdminMedicalHistoryController) CreateMedicalHistory(c *gin.Context) {
	var requestData dto.AdminCreateMedHistoryRequest
	if err := ginUtils.BindAndValidateBody(c, &requestData, ctlr.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	command := requestData.ToCommand()
	result := ctlr.bus.CommandBus.CreateMedicalHistory(*command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Created(c, result.ID, "Medical History")
}

func (ctrl AdminMedicalHistoryController) SoftDeleteMedicalHistory(c *gin.Context) {
	idInterface, err := ginUtils.ParseParamToEntityID(c, "medical_history")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "medical-history", c.Param("id")))
		return
	}

	mediHistID, valid := idInterface.(valueobject.MedHistoryID)
	if !valid {
		response.ServerError(c, httpError.InternalServerError(errors.New("invalid medical history ID type")))
		return
	}

	command := command.SoftDeleteMedHistCommand{
		ID:  mediHistID,
		CTX: c.Request.Context(),
	}

	result := ctrl.bus.CommandBus.SoftDeleteMedicalHistory(command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}
