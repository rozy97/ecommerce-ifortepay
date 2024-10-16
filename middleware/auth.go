package middleware

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (m *middleware) VerifyToken(c *fiber.Ctx) error {
	bearerToken := c.Get("Authorization")
	bearerTokenArr := strings.Split(bearerToken, " ")
	if strings.ToLower(bearerTokenArr[0]) != "bearer" || len(bearerTokenArr) != 2 {
		return c.SendStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(bearerTokenArr[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(m.env.SecretKey), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(map[string]string{"message": "invalid token"})
	}

	c.Locals("X-User-Id", token.Claims.(jwt.MapClaims)["user_id"])
	// log.Println(c.Locals("X-User-Id"))

	return c.Next()
}
