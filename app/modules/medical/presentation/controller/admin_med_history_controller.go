// Package controller defines the controllers for handling HTTP requests related to medical histories.
package controller

import (
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/query"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/infrastructure/bus"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/presentation/dto"

	httpError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/http"
	ginUtils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/gin_utils"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
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

	query := query.GetMedHistByIDQuery{
		ID:  mediHistID,
		CTX: c.Request.Context(),
	}

	medHistory, err := ctlr.bus.QueryBus.GetMedHistByID(query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Found(c, medHistory, "Medical History")
}

func (ctlr AdminMedicalHistoryController) CreateMedicalHistory(c *gin.Context) {
	var createData dto.AdminMedHistoryRequest
	if err := c.ShouldBindJSON(&createData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := ctlr.validator.Struct(createData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return

	}

	command := createData.ToCommand()
	result, err := ctlr.bus.CommandBus.CreateMedicalHistory(*command)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Created(c, result.ID, "Medical History")
}

func (ctlr AdminMedicalHistoryController) DeleteMedicalHistories(c *gin.Context) {
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

	softDelete := c.Query("is_soft_delete")
	if softDelete != "" && softDelete != "true" && softDelete != "false" {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}
	isSoftDelete := softDelete == "true"

	if err := ctlr.usecase.Delete(c.Request.Context(), mediHistID, isSoftDelete); err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.NoContent(c)
}
