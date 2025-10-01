package middleware

import (
	"api-e-ticketing/src/utils"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// RequireRole checks if user has the required role
func RequireRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.Error(c, fiber.StatusUnauthorized, "Missing token", nil, nil)
		}
		
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return utils.Error(c, fiber.StatusUnauthorized, "Invalid token format", nil, nil)
		}
		
		tokenStr := parts[1]
		secret := os.Getenv("JWT_SECRET")
		
		// Parse JWT
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			return utils.Error(c, fiber.StatusUnauthorized, "Invalid token", nil, nil)
		}
		
		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return utils.Error(c, fiber.StatusUnauthorized, "Invalid claims", nil, nil)
		}
		
		// Check role
		userRole, ok := claims["role"].(string)
		if !ok || userRole != role {
			return utils.Error(c, fiber.StatusUnauthorized, "Forbidden: insufficient role", nil, nil)
		}

		// Pass user info to next handler
		c.Locals("userRole", userRole)
		return c.Next()
	}
}
