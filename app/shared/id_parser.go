package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseID(c *gin.Context, param_name string) (uint, error) {
	idStr := c.Param("id")
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

	return uint(intValue), nil
}
