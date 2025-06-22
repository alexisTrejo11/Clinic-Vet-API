package petController

import (
	"context"

	petDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/dtos"
	petUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/usecase"
	utils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PetController struct {
	validator         *validator.Validate
	getPetByIdUseCase *petUsecase.GetPetByIdUseCase
	listPetsUseCase   *petUsecase.ListPetsUseCase
	createPetUseCae   *petUsecase.CreatePetUseCase
	updatePetUseCae   *petUsecase.UpdatePetUseCase
	deletePetUseCase  *petUsecase.DeletePetUseCase
}

func NewPetController(
	validator *validator.Validate,
	getPetByIdUseCase *petUsecase.GetPetByIdUseCase,
	listPetsUseCase *petUsecase.ListPetsUseCase,
	createPetUseCase *petUsecase.CreatePetUseCase,
	updatePetUseCase *petUsecase.UpdatePetUseCase,
	deletePetUseCase *petUsecase.DeletePetUseCase,

) *PetController {
	return &PetController{
		validator:         validator,
		getPetByIdUseCase: getPetByIdUseCase,
		listPetsUseCase:   listPetsUseCase,
		createPetUseCae:   createPetUseCase,
		updatePetUseCae:   updatePetUseCase,
		deletePetUseCase:  deletePetUseCase,
	}
}

func (c *PetController) ListPets(ctx *gin.Context) {
	pets, err := c.listPetsUseCase.Execute(context.TODO())
	if err != nil {
		apiResponse.BadRequest(ctx, err)
		return
	}

	apiResponse.Ok(ctx, pets)
}

func (c *PetController) GetPetById(ctx *gin.Context) {
	id, err := utils.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "pet_id", ctx.Param("id"))
		return
	}

	pet, err := c.getPetByIdUseCase.Execute(context.TODO(), id)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	apiResponse.Ok(ctx, pet)
}

func (c *PetController) CreatePet(ctx *gin.Context) {
	var petCreate petDTOs.PetCreate

	if err := ctx.ShouldBindBodyWithJSON(&petCreate); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.validator.Struct(&petCreate); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	pet, err := c.createPetUseCae.Execute(context.TODO(), petCreate)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	apiResponse.Created(ctx, pet)
}

func (c *PetController) UpdatePet(ctx *gin.Context) {
	id, err := utils.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "pet_id", ctx.Param("id"))
		return
	}

	var petCreate petDTOs.PetUpdate
	if err := ctx.ShouldBindBodyWithJSON(&petCreate); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.validator.Struct(&petCreate); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	pet, err := c.updatePetUseCae.Execute(context.TODO(), id, petCreate)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	apiResponse.Ok(ctx, pet)
}

func (c *PetController) DeletePet(ctx *gin.Context) {
	id, err := utils.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "pet_id", ctx.Param("id"))
		return
	}

	if err := c.deletePetUseCase.Execute(context.TODO(), id); err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	apiResponse.NoContent(ctx)
}
