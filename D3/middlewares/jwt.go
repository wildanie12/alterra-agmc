package middlewares

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const TOKEN = "THIS_IS_S3CR3T:V"

func VerifyJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return sendUnauthorizedResponse(c, "")
		}

		prefix := "Bearer "
		tokenStr := authHeader[len(prefix):]

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(TOKEN), nil
		})
		if err != nil {
			return sendUnauthorizedResponse(c, "error parsing token")
		}
		if !token.Valid {
			return sendUnauthorizedResponse(c, "inavlid token")
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		c.Set("auth_email", claims["email"])
		return next(c)
	}
}

func sendUnauthorizedResponse(c echo.Context, msg string) error {
	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"status":  "error",
		"code":    http.StatusUnauthorized,
		"message": map[bool]string{true: msg, false: "Unauthorized user"}[msg != ""],
	})
}
