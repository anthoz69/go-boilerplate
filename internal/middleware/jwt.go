package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/anthoz69/salepage-api/internal/core/domain"
	"github.com/anthoz69/salepage-api/shared/constants"
	"github.com/anthoz69/salepage-api/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// JWTAuth middleware for validating JWT tokens
func JWTAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the JWT token from the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.NewErrorResponse(c, constants.ErrCodeInvalidToken, "Missing authorization header")
		}

		// Check if the Authorization header has the Bearer prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return utils.NewErrorResponse(c, constants.ErrCodeInvalidToken, "Invalid authorization format")
		}

		tokenString := parts[1]

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Get the secret key from configuration
			secretKey := []byte(viper.GetString("jwt.secret"))
			return secretKey, nil
		})

		if err != nil {
			return utils.NewErrorResponse(c, constants.ErrCodeInvalidToken, fmt.Sprintf("Invalid token: %v", err))
		}

		// Check if the token is valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Store user information from claims in the context
			c.Locals("userID", claims["sub"])
			c.Locals("username", claims["username"])

			// Continue to the next middleware or route handler
			return c.Next()
		}

		return utils.NewErrorResponse(c, constants.ErrCodeInvalidToken, "Invalid token")
	}
}

// GenerateToken creates a new JWT token for a user
func GenerateToken(userID, username string, role domain.UserRole) (string, error) {
	// Create the claims
	claims := jwt.MapClaims{
		"sub":      userID,
		"username": username,
		"role":     role.String(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iat":      time.Now().Unix(),
	}

	// Create the token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	secretKey := []byte(viper.GetString("jwt.secret"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
