package med_hist_controller

import (
	"strconv"

	mhDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/application/dtos"
	medHistUsecases "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/application/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminMedicalHistoryController struct {
	usecases  *medHistUsecases.MedicalHistoryUseCase
	validator *validator.Validate
}

func NewAdminMedicalHistoryController(usecases *medHistUsecases.MedicalHistoryUseCase) *AdminMedicalHistoryController {
	return &AdminMedicalHistoryController{usecases: usecases, validator: validator.New()}
}

func (mhc AdminMedicalHistoryController) SearchMedicalHistories(c *gin.Context) {
	var serachParams mhDTOs.MedHistSearchParams
	if err := c.ShouldBindQuery(&serachParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := mhc.validator.Struct(serachParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	medHistories, err := mhc.usecases.Search(c.Request.Context(), serachParams)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": medHistories.Data, "metadata": gin.H{"pagintation": medHistories.Metadata}})
}

func (mhc AdminMedicalHistoryController) GetMedicalHistory(c *gin.Context) {
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

	medHistory, err := mhc.usecases.GetById(c.Request.Context(), idInt)
	if err != nil {
		c.JSON(404, gin.H{"error": "Medical history not found"})
		return
	}

	c.JSON(200, gin.H{"data": medHistory})
}

func (mhc AdminMedicalHistoryController) CreateMedicalHistories(c *gin.Context) {
	var createData mhDTOs.MedicalHistoryCreate
	if err := c.ShouldBindJSON(&createData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := mhc.validator.Struct(createData); err != nil {
		c.JSON(400, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	if err := mhc.usecases.Create(c.Request.Context(), createData); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create medical history", "details": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Medical history created successfully"})
}

func (mhc AdminMedicalHistoryController) UpdateMedicalHistories(c *gin.Context) {
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

	var updateData mhDTOs.MedicalHistoryUpdate
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := mhc.validator.Struct(updateData); err != nil {
		c.JSON(400, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	if err := mhc.usecases.Update(c.Request.Context(), idInt, updateData); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update medical history", "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Update Medical Histories"})
}

func (mhc AdminMedicalHistoryController) DeleteMedicalHistories(c *gin.Context) {
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

	if err := mhc.usecases.Delete(c.Request.Context(), idInt, isSoftDelete); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete medical history", "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Delete Medical Histories"})
}
