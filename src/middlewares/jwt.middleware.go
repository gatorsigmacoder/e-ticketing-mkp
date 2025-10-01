package middleware

import (
	"api-e-ticketing/src/database"
	"api-e-ticketing/src/models"
	"api-e-ticketing/src/utils"
	"os"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// jwtMiddleware validates JWT and user token version
func JwtMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return utils.Error(c, fiber.StatusUnauthorized, "Unauthorized access, please provide a valid token", nil, nil)
		},
		SuccessHandler: validateTokenAndUser,
	})
}

// validateTokenAndUser validates token and checks user token version
func validateTokenAndUser(c *fiber.Ctx) error {
	// Extract token from context
	userToken := c.Locals("user")
	token, ok := userToken.(*jwt.Token)
	if !ok {
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid token", nil, nil)
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid token", nil, nil)
	}

	// Get user ID and token version from claims
	userID, ok := claims["sub"].(string)
	if !ok {
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid token claims", nil, nil)
	}
	tokenVersion, ok := claims["token_version"].(float64)
	if !ok {
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid token claims", nil, nil)
	}

	// Validate token version against the database
	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, "User not found", nil, nil)
	}
	if user.TokenVersion != int(tokenVersion) {
		return utils.Error(c, fiber.StatusUnauthorized, "Token is invalid or expired", nil, nil)
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid token claims", nil, nil)
	}
	if int64(exp) < time.Now().Unix() {
		return utils.Error(c, fiber.StatusUnauthorized, "Token has expired", nil, nil)
	}

	return c.Next()
}
