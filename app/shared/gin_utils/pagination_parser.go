package ginUtils

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/gin-gonic/gin"
)

func GetPaginationParams(ctx gin.Context, orderBy string) page.PageData {
	pageNumber := ctx.Query("page_number")
}
