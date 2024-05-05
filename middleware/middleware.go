package middleware

import (
	"errors"
	"main/config"
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

// JWT token을 검증하는 미들웨어
func JWTMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("X-Authorization")

	result, err := JwtChecker(tokenString)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !result {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	return c.Next()
}

func JwtChecker(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(viper.GetString(config.JWT_SECRET)), nil
	})

	if err != nil {
		return false, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				return false, errors.New("invalid token")
			}
		}
	} else {
		return false, errors.New("invalid token")
	}

	return true, nil
}
