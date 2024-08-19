package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ParseID extracts and converts a path variable to an integer.
func ParseID(c *fiber.Ctx) (int32, error) {
	idStr := c.Params("id")
	if idStr == "" {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Id is Empty")
	}

	intValue, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Can't Process Id")
	}

	return int32(intValue), nil
}
