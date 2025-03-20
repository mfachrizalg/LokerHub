package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

// JWTAuth verifies that the request contains a valid JWT token
func JWTAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get cookie from request
		cookie := c.Cookies("LokerHubCookie")
		if cookie == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: No token provided",
			})
		}

		// Parse and validate the token
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			// Make sure the token method is what we expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Invalid token",
				"error":   err.Error(),
			})
		}

		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Token is not valid",
			})
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Invalid token claims",
			})
		}

		// Set user info in locals for later use in handlers
		c.Locals("userID", claims["id"])
		c.Locals("userEmail", claims["email"])
		c.Locals("userRole", claims["role"])

		return c.Next()
	}
}

// RoleAuth creates middleware that checks if the user has the required role
func RoleAuth(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the user role from the context (set by JWTAuth middleware)
		role := c.Locals("userRole")

		// If role is not set, user is not authenticated
		if role == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Authentication required",
			})
		}

		// Check if the user's role is allowed
		userRole, ok := role.(string)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error: Invalid role format",
			})
		}

		// Check if the role is in the allowed roles list
		allowed := false
		for _, r := range allowedRoles {
			if userRole == r {
				allowed = true
				break
			}
		}

		if !allowed {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden: Insufficient permissions",
			})
		}

		return c.Next()
	}
}
