package controller

import (
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/gin-gonic/gin"
)

func fetchVetParams(ctx *gin.Context) (dto.VetSearchParams, error) {
	var searchParams dto.VetSearchParams

	searchParams.PageSize = 10
	searchParams.SortDirection = page.DESC

	if err := ctx.ShouldBindQuery(&searchParams); err != nil {
		return dto.VetSearchParams{}, fmt.Errorf("parámetros de búsqueda inválidos")
	}

	if err := ctx.ShouldBindQuery(&searchParams.Filters); err != nil {
		return dto.VetSearchParams{}, fmt.Errorf("parámetros de filtro inválidos")
	}

	if searchParams.PageSize < 1 || searchParams.PageSize > 100 {
		return dto.VetSearchParams{}, fmt.Errorf("el límite debe estar entre 1 y 100")
	}

	if searchParams.PageNumber < 1 {
		return dto.VetSearchParams{}, fmt.Errorf("el numero de la pagina no puede ser menor a 1")
	}

	return searchParams, nil
}
