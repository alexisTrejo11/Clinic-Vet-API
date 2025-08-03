package vetController

import (
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	vetDtos "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
	"github.com/gin-gonic/gin"
)

func fetchVetParams(ctx *gin.Context) (vetDtos.VetSearchParams, error) {
	var searchParams vetDtos.VetSearchParams

	searchParams.Limit = 10
	searchParams.SortDirection = shared.DESC

	if err := ctx.ShouldBindQuery(&searchParams); err != nil {
		return vetDtos.VetSearchParams{}, fmt.Errorf("parámetros de búsqueda inválidos")
	}

	if err := ctx.ShouldBindQuery(&searchParams.Filters); err != nil {
		return vetDtos.VetSearchParams{}, fmt.Errorf("parámetros de filtro inválidos")
	}

	if searchParams.Limit < 1 || searchParams.Limit > 100 {
		return vetDtos.VetSearchParams{}, fmt.Errorf("el límite debe estar entre 1 y 100")
	}

	if searchParams.Offset < 0 {
		return vetDtos.VetSearchParams{}, fmt.Errorf("el offset no puede ser negativo")
	}

	return searchParams, nil
}
