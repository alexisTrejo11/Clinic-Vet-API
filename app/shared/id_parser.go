package shared

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseParamToInt(c *gin.Context, param_name string) (int, error) {
	idStr := c.Param(param_name)
	if idStr == "" {
		return 0, errors.New("empty id")
	}

	intValue, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("id must be a valid number")
	}

	if intValue < 0 {
		return 0, errors.New("ID cannot be negative")
	}

	return int(intValue), nil
}
