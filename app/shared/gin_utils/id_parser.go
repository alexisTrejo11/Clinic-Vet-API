// Package ginUtils contains all the helping operations to help controller to interact with gin framework
package ginUtils

import (
	"errors"
	"strconv"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/gin-gonic/gin"
)

func ParseParamToInt(c *gin.Context, paramName string) (int, error) {
	idStr := c.Param(paramName)
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

func ParseParamToEntityID(c *gin.Context, idParam string, entity string) (valueobject.IntegerID, error) {
	intValue, err := ParseParamToInt(c, idParam)
	if err != nil {
		return nil, err
	}

	entityID, err := valueobject.NewIDFactory(intValue, entity)
	if err != nil {
		return nil, err
	}

	return entityID, nil
}
