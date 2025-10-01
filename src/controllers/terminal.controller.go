package controllers

import (
	"api-e-ticketing/src/database"
	"api-e-ticketing/src/dtos"
	"api-e-ticketing/src/models"
	"api-e-ticketing/src/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateTerminal(c *fiber.Ctx) error {
	var terminalRequest dtos.TerminalRequest
	if err := c.BodyParser(&terminalRequest); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error(), nil, nil)
	}

	if err := utils.Validate.Struct(terminalRequest); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error(), nil, nil)
	}

	if terminalRequest.Name == "" {
		return utils.Error(c, fiber.StatusBadRequest, "Name was not included", nil, nil)
	}

	terminal := models.Terminal{
		Name: terminalRequest.Name,
	}

	if err := database.DB.Create(&terminal).Error; err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error(), nil, nil)
	}

	return utils.Success(c, fiber.StatusCreated, "Created", terminalRequest, nil)
}