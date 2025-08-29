package controller

import (
	"strconv"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/usecase"
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

func (ctlr AdminMedicalHistoryController) SearchMedicalHistories(ctx *gin.Context) {
	var serachParams dto.MedHistSearchParams
	if err := ctx.ShouldBindQuery(&serachParams); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ctlr.validator.Struct(serachParams); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	medHistories, err := ctlr.usecase.Search(ctx.Request.Context(), serachParams)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": medHistories.Data, "metadata": gin.H{"pagintation": medHistories.Metadata}})
}

func (ctlr AdminMedicalHistoryController) GetMedicalHistory(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID parameter is required"})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	medHistory, err := ctlr.usecase.GetByID(c.Request.Context(), idInt)
	if err != nil {
		c.JSON(404, gin.H{"error": "Medical history not found", "id": idInt, "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": medHistory})
}

func (ctlr AdminMedicalHistoryController) GetMedicalHistoryDetails(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID parameter is required"})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	medHistory, err := ctlr.usecase.GetByIDWithDeatils(c.Request.Context(), idInt)
	if err != nil {
		c.JSON(404, gin.H{"error": "Medical history not found", "id": idInt, "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": medHistory})
}

func (ctlr AdminMedicalHistoryController) CreateMedicalHistories(c *gin.Context) {
	var createData dto.MedicalHistoryCreate
	if err := c.ShouldBindJSON(&createData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := ctlr.validator.Struct(createData); err != nil {
		c.JSON(400, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	if err := ctlr.usecase.Create(c.Request.Context(), createData); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create medical history", "details": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Medical history created successfully"})
}

func (ctlr AdminMedicalHistoryController) UpdateMedicalHistories(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID parameter is required"})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var updateData dto.MedicalHistoryUpdate
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := ctlr.validator.Struct(updateData); err != nil {
		c.JSON(400, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	if err := ctlr.usecase.Update(c.Request.Context(), idInt, updateData); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update medical history", "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Update Medical Histories"})
}

func (ctlr AdminMedicalHistoryController) DeleteMedicalHistories(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID parameter is required"})
		return
	}

	softDelete := c.Query("is_soft_delete")
	if softDelete != "" && softDelete != "true" && softDelete != "false" {
		c.JSON(400, gin.H{"error": "Invalid is_soft_delete parameter"})
		return
	}
	isSoftDelete := softDelete == "true"

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := ctlr.usecase.Delete(c.Request.Context(), idInt, isSoftDelete); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete medical history", "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Delete Medical Histories"})
}
