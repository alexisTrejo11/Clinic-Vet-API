// Package controller defines the controllers for handling HTTP requests related to medical histories.
package controller

import (
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/usecase"
	httpError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/http"
	ginUtils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/gin_utils"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminMedicalHistoryController struct {
	usecase   *usecase.MedicalHistoryUseCase
	validator *validator.Validate
}

func NewAdminMedicalHistoryController(usecases *usecase.MedicalHistoryUseCase) *AdminMedicalHistoryController {
	return &AdminMedicalHistoryController{usecase: usecases, validator: validator.New()}
}

func (ctlr AdminMedicalHistoryController) SearchMedicalHistories(c *gin.Context) {
	var serachParams dto.MedHistSearchParams
	if err := c.ShouldBindQuery(&serachParams); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	if err := ctlr.validator.Struct(serachParams); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	medHistories, err := ctlr.usecase.Search(c.Request.Context(), serachParams)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.SuccessWithMeta(c, medHistories.Data, gin.H{"pagination": medHistories.Metadata})
}

func (ctlr AdminMedicalHistoryController) GetMedicalHistoryDetails(c *gin.Context) {
	idInterface, err := ginUtils.ParseParamToEntityID(c, "id", "medical_history")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "medical-history", c.Param("id")))
		return
	}

	mediHistID, valid := idInterface.(valueobject.MedHistoryID)
	if !valid {
		response.ServerError(c, httpError.InternalServerError(errors.New("invalid medical history ID type")))
		return
	}

	medHistory, err := ctlr.usecase.GetByIDWithDeatils(c.Request.Context(), mediHistID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, medHistory)
}

func (ctlr AdminMedicalHistoryController) CreateMedicalHistories(c *gin.Context) {
	var createData dto.MedicalHistoryCreate
	if err := c.ShouldBindJSON(&createData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := ctlr.validator.Struct(createData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return

	}

	if err := ctlr.usecase.Create(c.Request.Context(), createData); err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Created(c, "Medical history created successfully")
}

func (ctlr AdminMedicalHistoryController) DeleteMedicalHistories(c *gin.Context) {
	idInterface, err := ginUtils.ParseParamToEntityID(c, "id", "medical_history")
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
