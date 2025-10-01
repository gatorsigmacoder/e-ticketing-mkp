package controllers

import (
	"api-e-ticketing/src/database"
	"api-e-ticketing/src/dtos"
	"api-e-ticketing/src/models"
	"api-e-ticketing/src/utils"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func Login(ctx *fiber.Ctx) error {
	var req dtos.LoginRequest

	if err := ctx.BodyParser(&req); err != nil {
		return utils.Error(ctx, fiber.StatusBadRequest, "Invalid request body", nil, nil)
	}

	if req.Email == "" || req.Password == "" {
		return utils.Error(ctx, fiber.StatusBadRequest, "Password or email was not included", nil, nil)
	}

	var user models.User
	// Check by username or email
	result := database.DB.Debug().Where("username = ? OR email = ?", req.Email, req.Email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Handle case where user is not found
			return utils.Error(ctx, fiber.StatusUnauthorized, "Username/email or password is wrong", nil, nil)
		}
		return utils.Error(ctx, fiber.StatusInternalServerError, "Database error", result.Error.Error(), nil)
	}

	if !utils.CheckPassword(user.Password, req.Password) {
		return utils.Error(ctx, fiber.StatusUnauthorized, "Username or password is wrong", nil, nil)
	}

	claims := jwt.MapClaims{
		"sub":           user.ID,
		"eml":           user.Email,
		"role":           user.Role,
		"exp":           time.Now().Add(time.Hour * 12).Unix(),
		"token_version": user.TokenVersion,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	data := map[string]string{
		"token": t,
	}

	return utils.Success(ctx, fiber.StatusOK, "Success", data, nil)
}